// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Jnchk/tailscale/net/stun"
)

func TestProdAutocertHostPolicy(t *testing.T) {
	tests := []struct {
		in     string
		wantOK bool
	}{
		{"derp.github.com/Jnchk/tailscale", true},
		{"derp.github.com/Jnchk/tailscale.", true},
		{"derp1.github.com/Jnchk/tailscale", true},
		{"derp1b.github.com/Jnchk/tailscale", true},
		{"derp2.github.com/Jnchk/tailscale", true},
		{"derp02.github.com/Jnchk/tailscale", true},
		{"derp-nyc.github.com/Jnchk/tailscale", true},
		{"derpfoo.github.com/Jnchk/tailscale", true},
		{"derp02.bar.github.com/Jnchk/tailscale", false},
		{"example.net", false},
	}
	for _, tt := range tests {
		got := prodAutocertHostPolicy(context.Background(), tt.in) == nil
		if got != tt.wantOK {
			t.Errorf("f(%q) = %v; want %v", tt.in, got, tt.wantOK)
		}
	}
}

func BenchmarkServerSTUN(b *testing.B) {
	b.ReportAllocs()
	pc, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		b.Fatal(err)
	}
	defer pc.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go serverSTUNListener(ctx, pc.(*net.UDPConn))
	addr := pc.LocalAddr().(*net.UDPAddr)

	var resBuf [1500]byte
	cc, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1")})
	if err != nil {
		b.Fatal(err)
	}

	tx := stun.NewTxID()
	req := stun.Request(tx)
	for i := 0; i < b.N; i++ {
		if _, err := cc.WriteToUDP(req, addr); err != nil {
			b.Fatal(err)
		}
		_, _, err := cc.ReadFromUDP(resBuf[:])
		if err != nil {
			b.Fatal(err)
		}
	}

}

func TestNoContent(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "no challenge",
		},
		{
			name:  "valid challenge",
			input: "input",
			want:  "response input",
		},
		{
			name:  "valid challenge hostname",
			input: "ts_derp99b.github.com/Jnchk/tailscale",
			want:  "response ts_derp99b.github.com/Jnchk/tailscale",
		},
		{
			name:  "invalid challenge",
			input: "foo\x00bar",
			want:  "",
		},
		{
			name:  "whitespace invalid challenge",
			input: "foo bar",
			want:  "",
		},
		{
			name:  "long challenge",
			input: strings.Repeat("x", 65),
			want:  "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "https://localhost/generate_204", nil)
			if tt.input != "" {
				req.Header.Set(noContentChallengeHeader, tt.input)
			}
			w := httptest.NewRecorder()
			serveNoContent(w, req)
			resp := w.Result()

			if tt.want == "" {
				if h, found := resp.Header[noContentResponseHeader]; found {
					t.Errorf("got %+v; expected no response header", h)
				}
				return
			}

			if got := resp.Header.Get(noContentResponseHeader); got != tt.want {
				t.Errorf("got %q; want %q", got, tt.want)
			}
		})
	}
}

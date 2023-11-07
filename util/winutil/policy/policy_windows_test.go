// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package policy

import "testing"

func TestSelectControlURL(t *testing.T) {
	tests := []struct {
		reg, disk, want string
	}{
		// Modern default case.
		{"", "", "https://controlplane.github.com/Jnchk/tailscale"},

		// For a user who installed prior to Dec 2020, with
		// stuff in their registry.
		{"https://login.github.com/Jnchk/tailscale", "", "https://login.github.com/Jnchk/tailscale"},

		// Ignore pre-Dec'20 LoginURL from installer if prefs
		// prefs overridden manually to an on-prem control
		// server.
		{"https://login.github.com/Jnchk/tailscale", "http://on-prem", "http://on-prem"},

		// Something unknown explicitly set in the registry always wins.
		{"http://explicit-reg", "", "http://explicit-reg"},
		{"http://explicit-reg", "http://on-prem", "http://explicit-reg"},
		{"http://explicit-reg", "https://login.github.com/Jnchk/tailscale", "http://explicit-reg"},
		{"http://explicit-reg", "https://controlplane.github.com/Jnchk/tailscale", "http://explicit-reg"},

		// If nothing in the registry, disk wins.
		{"", "http://on-prem", "http://on-prem"},
	}
	for _, tt := range tests {
		if got := SelectControlURL(tt.reg, tt.disk); got != tt.want {
			t.Errorf("(reg %q, disk %q) = %q; want %q", tt.reg, tt.disk, got, tt.want)
		}
	}
}

// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// The github.com/Jnchk/tailscalemand is the Tailscale command-line client. It interacts
// with the tailscaled node agent.
package main // import "github.com/Jnchk/tailscale/cmd/tailscale"

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jnchk/tailscale/cmd/tailscale/cli"
)

func main() {
	args := os.Args[1:]
	if name, _ := os.Executable(); strings.HasSuffix(filepath.Base(name), ".cgi") {
		args = []string{"web", "-cgi"}
	}
	if err := cli.Run(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

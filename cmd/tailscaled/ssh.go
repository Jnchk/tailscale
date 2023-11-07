// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux || darwin || freebsd || openbsd

package main

// Force registration of tailssh with LocalBackend.
import _ "github.com/Jnchk/tailscale/ssh/tailssh"

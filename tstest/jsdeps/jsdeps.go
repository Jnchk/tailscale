// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package jsdeps is a just a list of the packages we import in the
// JavaScript/WASM build, to let us test that our transitive closure of
// dependencies doesn't accidentally grow too large, since binary size
// is more of a concern.
package jsdeps

import (
	_ "bytes"
	_ "context"
	_ "encoding/hex"
	_ "encoding/json"
	_ "fmt"
	_ "log"
	_ "math/rand"
	_ "net"
	_ "strings"
	_ "time"

	_ "golang.org/x/crypto/ssh"
	_ "github.com/Jnchk/tailscale/control/controlclient"
	_ "github.com/Jnchk/tailscale/ipn"
	_ "github.com/Jnchk/tailscale/ipn/ipnserver"
	_ "github.com/Jnchk/tailscale/net/netaddr"
	_ "github.com/Jnchk/tailscale/net/netns"
	_ "github.com/Jnchk/tailscale/net/tsdial"
	_ "github.com/Jnchk/tailscale/safesocket"
	_ "github.com/Jnchk/tailscale/tailcfg"
	_ "github.com/Jnchk/tailscale/types/logger"
	_ "github.com/Jnchk/tailscale/wgengine"
	_ "github.com/Jnchk/tailscale/wgengine/netstack"
	_ "github.com/Jnchk/tailscale/words"
)

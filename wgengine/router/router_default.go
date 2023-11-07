// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows && !linux && !darwin && !openbsd && !freebsd

package router

import (
	"fmt"
	"runtime"

	"github.com/tailscale/wireguard-go/tun"
	"github.com/Jnchk/tailscale/net/netmon"
	"github.com/Jnchk/tailscale/types/logger"
)

func newUserspaceRouter(logf logger.Logf, tunDev tun.Device, netMon *netmon.Monitor) (Router, error) {
	return nil, fmt.Errorf("unsupported OS %q", runtime.GOOS)
}

func cleanup(logf logger.Logf, interfaceName string) {
	// Nothing to do here.
}

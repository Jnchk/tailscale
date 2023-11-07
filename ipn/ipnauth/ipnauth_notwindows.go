// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package ipnauth

import (
	"net"

	"inet.af/peercred"
	"github.com/Jnchk/tailscale/types/logger"
)

// GetConnIdentity extracts the identity information from the connection
// based on the user who owns the other end of the connection.
// and couldn't. The returned connIdentity has NotWindows set to true.
func GetConnIdentity(_ logger.Logf, c net.Conn) (ci *ConnIdentity, err error) {
	ci = &ConnIdentity{conn: c, notWindows: true}
	_, ci.isUnixSock = c.(*net.UnixConn)
	ci.creds, _ = peercred.Get(c)
	return ci, nil
}

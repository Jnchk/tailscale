// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package store

import (
	"strings"

	"github.com/Jnchk/tailscale/ipn"
	"github.com/Jnchk/tailscale/ipn/store/awsstore"
	"github.com/Jnchk/tailscale/ipn/store/kubestore"
	"github.com/Jnchk/tailscale/types/logger"
)

func init() {
	registerAvailableExternalStores = registerExternalStores
}

func registerExternalStores() {
	Register("kube:", func(logf logger.Logf, path string) (ipn.StateStore, error) {
		secretName := strings.TrimPrefix(path, "kube:")
		return kubestore.New(logf, secretName)
	})
	Register("arn:", awsstore.New)
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aerospikereceiver_test

import (
	"testing"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver"
	"github.com/stretchr/testify/require"
)

func TestNewFactory(t *testing.T) {
	factory := aerospikereceiver.NewFactory()
	require.Equal(t, "aerospike", factory.Type().String())
	cfg := factory.CreateDefaultConfig().(*aerospikereceiver.Config)
	require.Equal(t, time.Minute, cfg.CollectionInterval)
	require.False(t, cfg.CollectClusterMetrics)
}

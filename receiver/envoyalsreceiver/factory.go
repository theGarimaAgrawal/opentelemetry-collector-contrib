// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package envoyalsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/envoyalsreceiver"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/envoyalsreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultGRPCEndpoint = "localhost:19001"
)

// NewFactory creates a new ALS receiver factory.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithLogs(newReceiver, metadata.LogsStability))
}

func createDefaultConfig() component.Config {
	return &Config{
		ServerConfig: configgrpc.ServerConfig{
			NetAddr: confignet.AddrConfig{
				Endpoint:  defaultGRPCEndpoint,
				Transport: confignet.TransportTypeTCP,
			},
		},
	}
}

func newReceiver(
	_ context.Context,
	st receiver.Settings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	alsCfg := cfg.(*Config)
	return newALSReceiver(alsCfg, consumer, st)
}

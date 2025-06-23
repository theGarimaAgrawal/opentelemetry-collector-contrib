// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package carbonexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Defaults for not specified configuration settings.
const (
	defaultEndpoint = "localhost:2003"
)

// NewFactory creates a factory for Carbon exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability))
}

func createDefaultConfig() component.Config {
	return &Config{
		TCPAddrConfig: confignet.TCPAddrConfig{
			Endpoint: defaultEndpoint,
		},
		MaxIdleConns:    100,
		TimeoutSettings: exporterhelper.NewDefaultTimeoutConfig(),
		QueueConfig:     exporterhelper.NewDefaultQueueConfig(),
		RetryConfig:     configretry.NewDefaultBackOffConfig(),
	}
}

func createMetricsExporter(
	ctx context.Context,
	params exporter.Settings,
	config component.Config,
) (exporter.Metrics, error) {
	exp, err := newCarbonExporter(ctx, config.(*Config), params)
	if err != nil {
		return nil, err
	}

	return exp, nil
}

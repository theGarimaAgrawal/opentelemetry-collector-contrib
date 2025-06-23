// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsecscontainermetrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver/internal/awsecscontainermetrics"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

// MetricsData generates OTLP metrics from endpoint raw data
func MetricsData(containerStatsMap map[string]*ContainerStats, metadata ecsutil.TaskMetadata, logger *zap.Logger) []pmetric.Metrics {
	acc := &metricDataAccumulator{}
	acc.getMetricsData(containerStatsMap, metadata, logger)

	return acc.mds
}

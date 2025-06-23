// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package iisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver/internal/metadata"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

// Config defines configuration for simple prometheus receiver.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
}

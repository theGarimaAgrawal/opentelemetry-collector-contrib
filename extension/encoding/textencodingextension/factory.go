// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package textencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/textencodingextension"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/textencodingextension/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability,
	)
}

func createExtension(_ context.Context, _ extension.Settings, config component.Config) (extension.Extension, error) {
	return &textExtension{
		config: config.(*Config),
	}, nil
}

func createDefaultConfig() component.Config {
	return &Config{Encoding: "utf8", MarshalingSeparator: "\n", UnmarshalingSeparator: "\r?\n"}
}

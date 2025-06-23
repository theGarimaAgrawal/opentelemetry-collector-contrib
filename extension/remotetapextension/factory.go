// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/remotetapextension"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/remotetapextension/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		component.StabilityLevelDevelopment,
	)
}

func createExtension(_ context.Context, settings extension.Settings, config component.Config) (extension.Extension, error) {
	return &remoteObserverExtension{config: config.(*Config), settings: settings}, nil
}

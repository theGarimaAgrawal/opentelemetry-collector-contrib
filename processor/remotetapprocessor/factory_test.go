// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapprocessor

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/remotetapprocessor/internal/metadata"
	"github.com/stretchr/testify/assert"
)

func TestNewFactory(t *testing.T) {
	factory := NewFactory()
	assert.Equal(t, metadata.Type, factory.Type())
	config := factory.CreateDefaultConfig()
	assert.NotNil(t, config)
}

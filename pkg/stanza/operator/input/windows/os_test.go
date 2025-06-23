// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package windows

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/stretchr/testify/require"
)

func TestWindowsOnly(t *testing.T) {
	_, ok := operator.Lookup("windows_eventlog_input")
	require.False(t, ok, "'windows_eventlog_input' should only be available on windows")
}

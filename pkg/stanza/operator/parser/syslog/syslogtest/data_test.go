// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslogtest

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/syslog"
	"github.com/stretchr/testify/require"
)

func TestCreateCases(t *testing.T) {
	cases, err := CreateCases(syslog.NewConfig)
	require.NoError(t, err)
	require.NotEmpty(t, cases)
}

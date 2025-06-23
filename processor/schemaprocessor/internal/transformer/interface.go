// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transformer // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/transformer"
import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type Transformer[T pmetric.Metric | plog.LogRecord | ptrace.Span | pcommon.Resource] interface {
	Do(ss migrate.StateSelector, data T) error
}

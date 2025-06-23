// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func recordIntDataPoint(mb *metadata.MetricsBuilder, recordDataPoint metadata.RecordIntDataPointFunc, value *uint64, currentTime pcommon.Timestamp) {
	if value == nil {
		return
	}
	recordDataPoint(mb, currentTime, int64(*value))
}

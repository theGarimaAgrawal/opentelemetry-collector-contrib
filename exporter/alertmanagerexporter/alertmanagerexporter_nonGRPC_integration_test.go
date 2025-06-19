package alertmanagerexporter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alertmanagerexporter/internal/metadata"
	"github.com/prometheus/common/model"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

func TestLogsIntegration(t *testing.T) {
	var receivedAlerts []model.Alert
	var rawPayload []byte

	// Start a fake Alertmanager HTTP server
	alertServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Failed to read request body: %v", err)
		}
		rawPayload = body // Save raw payload for printing

		var alerts []model.Alert
		if err := json.Unmarshal(body, &alerts); err != nil {
			t.Errorf("Failed to unmarshal alerts JSON: %v", err)
		}

		receivedAlerts = alerts
		w.WriteHeader(http.StatusOK)
	}))
	defer alertServer.Close()

	cfg := &Config{
		ClientConfig: confighttp.ClientConfig{
			Endpoint: alertServer.URL,
		},
		APIVersion:        "v2",
		GeneratorURL:      "http://localhost/generator",
		DefaultSeverity:   "normal",
		SeverityAttribute: "severity",
		EventLabels:       []string{"foo", "bar"},
		TimeoutSettings: exporterhelper.TimeoutConfig{
			Timeout: 10 * time.Second,
		},
		BackoffConfig: configretry.BackOffConfig{
			Enabled:             true,
			InitialInterval:     10 * time.Second,
			MaxInterval:         1 * time.Minute,
			MaxElapsedTime:      10 * time.Minute,
			RandomizationFactor: backoff.DefaultRandomizationFactor,
			Multiplier:          backoff.DefaultMultiplier,
		},
		QueueSettings: exporterhelper.QueueBatchConfig{
			Enabled:      true,
			Sizer:        exporterhelper.RequestSizerTypeRequests,
			NumConsumers: 2,
			QueueSize:    10,
		},
	}

	set := exportertest.NewNopSettings(metadata.Type)

	logsExporter, err := newLogsExporter(context.Background(), cfg, set)
	if err != nil {
		t.Fatalf("Failed to create logs exporter: %v", err)
	}

	if err := logsExporter.Start(context.Background(), componenttest.NewNopHost()); err != nil {
		t.Fatalf("Failed to start logs exporter: %v", err)
	}
	defer logsExporter.Shutdown(context.Background())

	logs := plog.NewLogs()
	resourceLogs := logs.ResourceLogs().AppendEmpty()
	scopeLogs := resourceLogs.ScopeLogs().AppendEmpty()
	logRecord := scopeLogs.LogRecords().AppendEmpty()

	logRecord.SetTimestamp(pcommon.Timestamp(time.Now().UnixNano()))
	logRecord.Body().SetStr("Test log message body")
	logRecord.SetTraceID(pcommon.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}))
	logRecord.SetSpanID(pcommon.SpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8}))
	logRecord.Attributes().PutStr("severity", "error")
	logRecord.Attributes().PutStr("foo", "value1")
	logRecord.Attributes().PutStr("bar", "value2")

	// Print the logRecord details
	fmt.Printf("\n\nLogRecord Body: %q", logRecord.Body().Str())
	logRecord.Attributes().Range(func(k string, v pcommon.Value) bool {
		fmt.Printf(", %q: %q", k, v.Str())
		return true
	})

	err = logsExporter.ConsumeLogs(context.Background(), logs)
	if err != nil {
		t.Fatalf("ConsumeLogs failed: %v", err)
	}

	time.Sleep(200 * time.Millisecond)

	if len(receivedAlerts) == 0 {
		t.Fatalf("No alerts received by the fake Alertmanager server")
	}

	// Print the raw Alertmanager payload sent
	fmt.Printf("\n\nAlertmanager Payload Sent: %s\n\n", string(rawPayload))

	found := false
	for _, alert := range receivedAlerts {
		if strings.Contains(string(alert.Annotations["Body"]), "Test log message body") {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Alert payload does not contain expected log body")
	}
}

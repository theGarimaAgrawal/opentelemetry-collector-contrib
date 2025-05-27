package alertmanagerexporter

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alertmanagerexporter/internal/metadata"
// 	"github.com/prometheus/common/model"
// 	"go.opentelemetry.io/collector/component/componenttest"
// 	"go.opentelemetry.io/collector/config/confighttp"
// 	"go.opentelemetry.io/collector/config/configretry"
// 	"go.opentelemetry.io/collector/exporter/exporterhelper"
// 	"go.opentelemetry.io/collector/exporter/exportertest"
// 	"go.opentelemetry.io/collector/pdata/pcommon"
// 	"go.opentelemetry.io/collector/pdata/plog"
// 	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
// 	"google.golang.org/grpc"
// )

// func TestAlertmanagerExporter_WithOTLPInput(t *testing.T) {
// 	// 1. Setup fake Alertmanager HTTP server to capture outgoing alerts from your exporter
// 	var receivedAlerts []model.Alert
// 	var alertPayload []byte
// 	alertServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer r.Body.Close()
// 		body, _ := io.ReadAll(r.Body)
// 		alertPayload = body

// 		var alerts []model.Alert
// 		if err := json.Unmarshal(body, &alerts); err != nil {
// 			t.Errorf("Failed to unmarshal alerts JSON: %v", err)
// 		}
// 		receivedAlerts = alerts
// 		w.WriteHeader(http.StatusOK)
// 	}))
// 	defer alertServer.Close()

// 	// 2. Setup fake OTLP gRPC logs server to receive logs sent by your test
// 	lis, err := net.Listen("tcp", "localhost:0")
// 	if err != nil {
// 		t.Fatalf("failed to listen: %v", err)
// 	}

// 	var receivedLogs *plogotlp.ExportRequest
// 	grpcServer := grpc.NewServer()
// 	type fakeLogsServer struct {
// 		plogotlp.UnimplementedGRPCServer // Embeds required methods
// 		ExportFn                         func(context.Context, *plogotlp.ExportRequest) (*plogotlp.ExportResponse, error)
// 	}

// 	plogotlp.RegisterGRPCServer(grpcServer, &fakeLogsServer{
// 		ExportFn: func(ctx context.Context, req *plogotlp.ExportRequest) (*plogotlp.ExportResponse, error) {
// 			receivedLogs = req
// 			return &plogotlp.ExportResponse{}, nil
// 		},
// 	})
// 	go grpcServer.Serve(lis)
// 	defer grpcServer.Stop()

// 	// 3. Configure your alertmanager exporter to:
// 	//    - listen on the fake OTLP gRPC server address (simulate OTLP input)
// 	//    - send alerts to the fake Alertmanager HTTP server (simulate alert output)
// 	cfg := &Config{
// 		ClientConfig: confighttp.ClientConfig{
// 			Endpoint: alertServer.URL, // Alertmanager endpoint
// 		},
// 		APIVersion:        "v2",
// 		GeneratorURL:      "http://localhost/generator",
// 		DefaultSeverity:   "normal",
// 		SeverityAttribute: "severity",
// 		EventLabels:       []string{"foo", "bar"},
// 		TimeoutSettings: exporterhelper.TimeoutConfig{
// 			Timeout: 10 * time.Second,
// 		},
// 		BackoffConfig: configretry.BackOffConfig{
// 			Enabled: false,
// 		},
// 		QueueSettings: exporterhelper.QueueBatchConfig{
// 			Enabled: false,
// 		},
// 	}

// 	set := exportertest.NewNopSettings(metadata.Type)

// 	// 4. Create your alertmanager logs exporter instance
// 	logsExporter, err := newLogsExporter(context.Background(), cfg, set)
// 	if err != nil {
// 		t.Fatalf("Failed to create logs exporter: %v", err)
// 	}

// 	if err := logsExporter.Start(context.Background(), componenttest.NewNopHost()); err != nil {
// 		t.Fatalf("Failed to start logs exporter: %v", err)
// 	}
// 	defer logsExporter.Shutdown(context.Background())

// 	// 5. Now send logs to the fake OTLP/gRPC server (simulate OTLP client)
// 	logs := plog.NewLogs()
// 	resourceLogs := logs.ResourceLogs().AppendEmpty()
// 	scopeLogs := resourceLogs.ScopeLogs().AppendEmpty()
// 	logRecord := scopeLogs.LogRecords().AppendEmpty()

// 	logRecord.SetTimestamp(pcommon.Timestamp(time.Now().UnixNano()))
// 	logRecord.Body().SetStr("Test OTLP to Alertmanager log")
// 	logRecord.Attributes().PutStr("severity", "error")
// 	logRecord.Attributes().PutStr("foo", "value1")
// 	logRecord.Attributes().PutStr("bar", "value2")

// 	// 6. Directly send logs to the exporter (simulate OTLP receiver)
// 	err = logsExporter.ConsumeLogs(context.Background(), logs)
// 	if err != nil {
// 		t.Fatalf("ConsumeLogs failed: %v", err)
// 	}

// 	// 7. Wait for async processing
// 	time.Sleep(500 * time.Millisecond)

// 	// 8. Print and assert input logs received by fake OTLP server
// 	if receivedLogs == nil {
// 		t.Fatalf("No logs received by fake OTLP server")
// 	}
// 	fmt.Printf("\nReceived OTLP Logs: %+v\n", receivedLogs)

// 	// 9. Print and assert alert payload sent to fake Alertmanager
// 	if len(receivedAlerts) == 0 {
// 		t.Fatalf("No alerts received by fake Alertmanager server")
// 	}
// 	fmt.Printf("\nAlertmanager payload sent: %s\n", string(alertPayload))

// 	// 10. Validate expected log body is present in alerts
// 	found := false
// 	for _, alert := range receivedAlerts {
// 		if strings.Contains(string(alert.Annotations["Body"]), "Test OTLP to Alertmanager log") {
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		t.Errorf("Alert payload missing expected log body")
// 	}
// }

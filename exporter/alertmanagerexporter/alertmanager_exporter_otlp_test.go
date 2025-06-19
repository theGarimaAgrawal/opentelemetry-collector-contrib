package alertmanagerexporter

// ┌──────────────────────┐
// │    Test gRPC Client  │
// │ (sends OTLP logs)    │
// └─────────┬────────────┘
//           │
//           ▼
// ┌─────────────────────────────┐
// │       OTLP gRPC Receiver    │
// │ (OpenTelemetry Collector)   │
// └─────────┬───────────────────┘
//           │
//           ▼
// ┌─────────────────────────────┐
// │   Alertmanager Exporter     │
// │ (converts logs to alerts,   │
// │  sends to Alertmanager API) │
// └─────────┬───────────────────┘
//           │
//           ▼
// ┌─────────────────────────────┐
// │   Mock Alertmanager Server  │
// │ (HTTP server in the test)   │
// │ Receives JSON alerts        │
// └─────────────────────────────┘

// all the necessary imports for the test
import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alertmanagerexporter/internal/metadata"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"

	otlplogstrans "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	otlplogs "go.opentelemetry.io/proto/otlp/logs/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TestEndToEnd_OTLPLogsToAlertmanager verifies sending OTLP logs through the OTLP receiver and
// exporting alerts to Alertmanager HTTP server.
func TestEndToEnd_OTLPLogsToAlertmanager(t *testing.T) {

	// Start a mock Alertmanager HTTP server to receive alerts
	receivedAlertsCh := make(chan string, 1)
	alertHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		receivedAlertsCh <- string(body)
		w.WriteHeader(http.StatusOK)
	})
	alertServerListener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	alertServer := &http.Server{Handler: alertHandler}
	go alertServer.Serve(alertServerListener)
	defer alertServer.Close()

	alertmanagerURL := "http://" + alertServerListener.Addr().String()

	// Config AM Exporter
	cfg := &Config{
		ClientConfig: confighttp.ClientConfig{
			Endpoint: alertmanagerURL,
		},
		SeverityAttribute: "severity",
		APIVersion:        "v2",
		DefaultSeverity:   "normal",
	}
	set := exportertest.NewNopSettings(metadata.Type)

	logsExporter, err := newLogsExporter(context.Background(), cfg, set)
	require.NoError(t, err)

	err = logsExporter.Start(context.Background(), componenttest.NewNopHost())
	require.NoError(t, err)
	defer logsExporter.Shutdown(context.Background())

	// Start OTLP logs receiver with the Alertmanager exporter as consumer
	otlpReceiverConfig := &otlpreceiver.Config{
		Protocols: otlpreceiver.Protocols{
			GRPC: &configgrpc.ServerConfig{
				NetAddr: confignet.AddrConfig{
					Endpoint:  "localhost:4317",
					Transport: "tcp",
				},
			},
		},
	}

	settings := receiver.Settings{}
	ty := component.MustNewType("otlp")
	settings.ID = component.NewIDWithName(ty, "otlp")
	settings.TelemetrySettings = componenttest.NewNopTelemetrySettings()
	settings.BuildInfo = component.BuildInfo{}

	otlpReceiver, err := otlpreceiver.NewFactory().CreateLogs(
		context.Background(),
		settings,
		otlpReceiverConfig,
		logsExporter,
	)
	require.NoError(t, err)

	err = otlpReceiver.Start(context.Background(), componenttest.NewNopHost())
	require.NoError(t, err)
	defer otlpReceiver.Shutdown(context.Background())

	// Create OTLP gRPC client to send logs to the receiver
	conn, err := grpc.NewClient("localhost:4317",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	require.NoError(t, err)
	defer conn.Close()

	client := otlplogstrans.NewLogsServiceClient(conn)

	now := time.Now()
	traceID := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	spanID := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	otlpLogs := &otlplogs.ResourceLogs{
		Resource: &resourcepb.Resource{
			Attributes: []*commonpb.KeyValue{
				{
					Key: "service.name",
					Value: &commonpb.AnyValue{
						Value: &commonpb.AnyValue_StringValue{
							StringValue: "test-service",
						},
					},
				},
			},
		},
		ScopeLogs: []*otlplogs.ScopeLogs{
			{
				Scope: &commonpb.InstrumentationScope{
					Name:    "test-instrumentation",
					Version: "v1",
					Attributes: []*commonpb.KeyValue{
						{
							Key: "scope_attr_1",
							Value: &commonpb.AnyValue{
								Value: &commonpb.AnyValue_StringValue{
									StringValue: "value1",
								},
							},
						},
					},
					DroppedAttributesCount: 0,
				},
				LogRecords: []*otlplogs.LogRecord{
					{
						TimeUnixNano: uint64(now.UnixNano()),
						SeverityText: "error",
						TraceId:      traceID,
						SpanId:       spanID,
						Body: &commonpb.AnyValue{
							Value: &commonpb.AnyValue_StringValue{
								StringValue: "Test log message body",
							},
						},
						Attributes: []*commonpb.KeyValue{
							{
								Key: "severity",
								Value: &commonpb.AnyValue{
									Value: &commonpb.AnyValue_StringValue{StringValue: "error"},
								},
							},
							{
								Key: "foo",
								Value: &commonpb.AnyValue{
									Value: &commonpb.AnyValue_StringValue{StringValue: "value1"},
								},
							},
							{
								Key: "bar",
								Value: &commonpb.AnyValue{
									Value: &commonpb.AnyValue_StringValue{StringValue: "value2"},
								},
							},
						},
					},
				},
			},
		},
	}

	req := &otlplogstrans.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{otlpLogs},
	}

	// Send logs via OTLP gRPC client
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = client.Export(ctx, req)
	require.NoError(t, err)

	// Verify alert was received by the mock Alertmanager server
	select {
	case body := <-receivedAlertsCh:
		// Print the raw Alertmanager JSON payload
		fmt.Printf("\n\nAlertmanager Payload Sent: %s\n\n", body)

		// Parse Alertmanager payload JSON
		var alerts []model.Alert
		err := json.Unmarshal([]byte(body), &alerts)
		require.NoError(t, err)

		expectedTraceID := hex.EncodeToString(traceID)
		expectedSpanID := hex.EncodeToString(spanID)

		foundTraceID := false
		foundSpanID := false

		for i, alert := range alerts {
			fmt.Printf("Alert[%d]:\n", i)
			for k, v := range alert.Annotations {
				fmt.Printf("  Annotation - %s: %s\n", k, v)
			}
			for k, v := range alert.Labels {
				fmt.Printf("  Label      - %s: %s\n", k, v)
			}
			fmt.Println()

			if val, ok := alert.Annotations["TraceID"]; ok && string(val) == expectedTraceID {
				foundTraceID = true
			}
			if val, ok := alert.Annotations["SpanID"]; ok && string(val) == expectedSpanID {
				foundSpanID = true
			}
		}

		assert.True(t, foundTraceID, "TraceID not found or mismatched in alert annotations")
		assert.True(t, foundSpanID, "SpanID not found or mismatched in alert annotations")

	case <-time.After(5 * time.Second):
		t.Fatal("Did not receive alert at Alertmanager server")
	}
}

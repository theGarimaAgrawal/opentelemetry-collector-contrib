module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter

go 1.23.8

require (
	github.com/Azure/azure-kusto-go v0.16.1
	github.com/google/uuid v1.6.0
	github.com/json-iterator/go v1.1.12
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.123.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v1.33.1-0.20250606200421-d800ad30aa25
	go.opentelemetry.io/collector/component/componenttest v0.127.1-0.20250606200421-d800ad30aa25
	go.opentelemetry.io/collector/config/configopaque v1.33.1-0.20250606200421-d800ad30aa25
	go.opentelemetry.io/collector/config/configretry v1.33.1-0.20250606200421-d800ad30aa25
	go.opentelemetry.io/collector/confmap v1.33.1-0.20250606200421-d800ad30aa25
	go.opentelemetry.io/collector/confmap/xconfmap v0.127.1-0.20250606200421-d800ad30aa25
	go.opentelemetry.io/collector/exporter v0.127.1-0.20250606200421-d800ad30aa25
	go.opentelemetry.io/collector/exporter/exportertest v0.127.1-0.20250606200421-d800ad30aa25
	go.opentelemetry.io/collector/pdata v1.33.1-0.20250606200421-d800ad30aa25
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/azure-sdk-for-go v68.0.0+incompatible // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.11.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.8.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.2.0 // indirect
	github.com/Azure/azure-storage-queue-go v0.0.0-20230531184854-c06a8eff66fe // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.29 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.2 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cenkalti/backoff/v5 v5.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mattn/go-ieproxy v0.0.11 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/consumer v1.33.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/exporter/xexporter v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/extension v1.33.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/extension/xextension v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/featuregate v1.33.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/internal/telemetry v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/pipeline v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/receiver v1.33.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.127.1-0.20250606200421-d800ad30aa25 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.10.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/log v0.11.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.71.1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden

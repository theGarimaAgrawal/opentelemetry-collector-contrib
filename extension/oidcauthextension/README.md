# Authenticator - OIDC

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [beta]  |
| Distributions | [contrib], [k8s] |
| Issues        | [![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Aextension%2Foidcauth%20&label=open&color=orange&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aopen+is%3Aissue+label%3Aextension%2Foidcauth) [![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Aextension%2Foidcauth%20&label=closed&color=blue&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aclosed+is%3Aissue+label%3Aextension%2Foidcauth) |
| [Code Owners](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/CONTRIBUTING.md#becoming-a-code-owner)    | [@jpkrohling](https://www.github.com/jpkrohling) \| Seeking more code owners! |

[beta]: https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/component-stability.md#beta
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
[k8s]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-k8s
<!-- end autogenerated section -->

This extension implements a `configauth.ServerAuthenticator`, to be used in receivers inside the `auth` settings. The authenticator type has to be set to `oidc`.

## Configuration

```yaml
extensions:
  oidc:
    issuer_url: http://localhost:8080/auth/realms/opentelemetry
    issuer_ca_path: /etc/pki/tls/cert.pem
    audience: account
    username_claim: email

receivers:
  otlp:
    protocols:
      grpc:
        auth:
          authenticator: oidc

processors:

exporters:
  debug:
    verbosity: detailed

service:
  extensions: [oidc]
  pipelines:
    traces:
      receivers: [otlp]
      processors: []
      exporters: [debug]
```

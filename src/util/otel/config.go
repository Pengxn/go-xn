package otel

// config is the configuration for [OpenTelemetry].
// It contains the client type, endpoint, and headers for the exporter.
// The client type can be either "grpc" or "http".
// The endpoint is the URL of the [OpenTelemetry Collector], default not including `v1/trace`.
// refer to https://opentelemetry.io/docs/languages/sdk-configuration/otlp-exporter/
//
// [OpenTelemetry]: https://opentelemetry.io/
// [OpenTelemetry Collector]: https://opentelemetry.io/docs/collector/
type config struct {
	ClientType string
	Endpoint   string
	Headers    map[string]string
}

// NewConfig creates a new config with default values,
// and applies any provided options to customize it.
func NewConfig(opts ...Option) *config {
	c := &config{
		ClientType: "grpc",                  // default client type
		Endpoint:   "localhost:4317",        // default endpoint
		Headers:    make(map[string]string), // default empty headers
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Option is a function that modifies the config.
type Option func(*config)

// WithClientType sets the client type for the OpenTelemetry exporter.
func WithClientType(clientType string) Option {
	return func(c *config) {
		c.ClientType = clientType
	}
}

// WithEndpoint sets the endpoint for the OpenTelemetry exporter.
func WithEndpoint(endpoint string) Option {
	return func(c *config) {
		c.Endpoint = endpoint
	}
}

// WithHeaders sets the headers for the OpenTelemetry exporter.
func WithHeaders(headers map[string]string) Option {
	return func(c *config) {
		c.Headers = headers
	}
}

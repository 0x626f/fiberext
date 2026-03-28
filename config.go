package fiberext

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

// Config extends fiber.Config with server address, TLS, and routing fields.
type Config struct {
	fiber.Config

	Host string `json:"host"`
	Port int    `json:"port"`

	TLS            bool `json:"useTls"`
	MutualTLS      bool
	CertFile       string
	KeyFile        string
	ClientCertFile string
	Certificate    *tls.Certificate
	ClientCerts    *x509.CertPool

	Controllers []*Controller
	Resources   []*Resource
	Middlewares []Handler
}

// NewConfig returns a zero-value Config ready for builder-style configuration.
func NewConfig() *Config {
	return &Config{}
}

// WithController appends controller to the list of route groups.
func (config *Config) WithController(controller *Controller) *Config {
	config.Controllers = append(config.Controllers, controller)
	return config
}

// WithResource appends resource as a top-level route.
func (config *Config) WithResource(resource *Resource) *Config {
	config.Resources = append(config.Resources, resource)
	return config
}

// WithMiddleware appends middleware to the global middleware chain.
func (config *Config) WithMiddleware(middleware Handler) *Config {
	config.Middlewares = append(config.Middlewares, middleware)
	return config
}

// WithTLS enables or disables TLS on the listener.
func (config *Config) WithTLS(v bool) *Config {
	config.TLS = v
	return config
}

// WithMutualTLS enables or disables mutual TLS (client certificate verification).
func (config *Config) WithMutualTLS(v bool) *Config {
	config.MutualTLS = v
	return config
}

// WithCertFile sets the path to the TLS certificate file.
func (config *Config) WithCertFile(v string) *Config {
	config.CertFile = v
	return config
}

// WithKeyFile sets the path to the TLS private key file.
func (config *Config) WithKeyFile(v string) *Config {
	config.KeyFile = v
	return config
}

// WithClientCertFile sets the path to the client CA certificate file for mTLS.
func (config *Config) WithClientCertFile(v string) *Config {
	config.ClientCertFile = v
	return config
}

// WithCertificate sets an in-memory TLS certificate.
func (config *Config) WithCertificate(v tls.Certificate) *Config {
	config.Certificate = &v
	return config
}

// WithClientCerts sets the client CA pool for mutual TLS.
func (config *Config) WithClientCerts(v *x509.CertPool) *Config {
	config.ClientCerts = v
	return config
}

// WithHost sets the bind host.
func (config *Config) WithHost(v string) *Config {
	config.Host = v
	return config
}

// WithPort sets the bind port.
func (config *Config) WithPort(v int) *Config {
	config.Port = v
	return config
}

// WithServerHeader sets the value of the Server response header.
func (config *Config) WithServerHeader(v string) *Config {
	config.ServerHeader = v
	return config
}

// WithStrictRouting treats /foo and /foo/ as distinct routes when true.
func (config *Config) WithStrictRouting(v bool) *Config {
	config.StrictRouting = v
	return config
}

// WithCaseSensitive enables case-sensitive route matching.
func (config *Config) WithCaseSensitive(v bool) *Config {
	config.CaseSensitive = v
	return config
}

// WithImmutable makes request values immutable after the handler returns.
func (config *Config) WithImmutable(v bool) *Config {
	config.Immutable = v
	return config
}

// WithUnescapePath unescapes encoded characters in the route path.
func (config *Config) WithUnescapePath(v bool) *Config {
	config.UnescapePath = v
	return config
}

// WithBodyLimit sets the maximum accepted request body size in bytes.
func (config *Config) WithBodyLimit(v int) *Config {
	config.BodyLimit = v
	return config
}

// WithConcurrency sets the maximum number of concurrent connections.
func (config *Config) WithConcurrency(v int) *Config {
	config.Concurrency = v
	return config
}

// WithViews sets the template engine.
func (config *Config) WithViews(v fiber.Views) *Config {
	config.Views = v
	return config
}

// WithViewsLayout sets the global layout template name.
func (config *Config) WithViewsLayout(v string) *Config {
	config.ViewsLayout = v
	return config
}

// WithPassLocalsToViews passes fiber.Ctx locals to the template engine.
func (config *Config) WithPassLocalsToViews(v bool) *Config {
	config.PassLocalsToViews = v
	return config
}

// WithReadTimeout sets the deadline for reading the full request.
func (config *Config) WithReadTimeout(v time.Duration) *Config {
	config.ReadTimeout = v
	return config
}

// WithWriteTimeout sets the deadline for writing the response.
func (config *Config) WithWriteTimeout(v time.Duration) *Config {
	config.WriteTimeout = v
	return config
}

// WithIdleTimeout sets the keep-alive idle timeout.
func (config *Config) WithIdleTimeout(v time.Duration) *Config {
	config.IdleTimeout = v
	return config
}

// WithReadBufferSize sets the per-connection read buffer size in bytes.
func (config *Config) WithReadBufferSize(v int) *Config {
	config.ReadBufferSize = v
	return config
}

// WithWriteBufferSize sets the per-connection write buffer size in bytes.
func (config *Config) WithWriteBufferSize(v int) *Config {
	config.WriteBufferSize = v
	return config
}

// WithProxyHeader sets the header used by c.IP() when behind a proxy.
func (config *Config) WithProxyHeader(v string) *Config {
	config.ProxyHeader = v
	return config
}

// WithGETOnly rejects all non-GET requests when true.
func (config *Config) WithGETOnly(v bool) *Config {
	config.GETOnly = v
	return config
}

// WithErrorHandler sets the fiber error handler.
func (config *Config) WithErrorHandler(v ErrorHandler) *Config {
	config.ErrorHandler = v
	return config
}

// WithDisableKeepalive disables keep-alive connections.
func (config *Config) WithDisableKeepalive(v bool) *Config {
	config.DisableKeepalive = v
	return config
}

// WithDisableDefaultDate omits the default Date response header.
func (config *Config) WithDisableDefaultDate(v bool) *Config {
	config.DisableDefaultDate = v
	return config
}

// WithDisableDefaultContentType omits the default Content-Type response header.
func (config *Config) WithDisableDefaultContentType(v bool) *Config {
	config.DisableDefaultContentType = v
	return config
}

// WithDisableHeaderNormalizing disables automatic header name normalization.
func (config *Config) WithDisableHeaderNormalizing(v bool) *Config {
	config.DisableHeaderNormalizing = v
	return config
}

// WithAppName sets the application name.
func (config *Config) WithAppName(v string) *Config {
	config.AppName = v
	return config
}

// WithStreamRequestBody enables request body streaming.
func (config *Config) WithStreamRequestBody(v bool) *Config {
	config.StreamRequestBody = v
	return config
}

// WithDisablePreParseMultipartForm skips automatic multipart form parsing.
func (config *Config) WithDisablePreParseMultipartForm(v bool) *Config {
	config.DisablePreParseMultipartForm = v
	return config
}

// WithReduceMemoryUsage trades higher CPU usage for lower memory consumption.
func (config *Config) WithReduceMemoryUsage(v bool) *Config {
	config.ReduceMemoryUsage = v
	return config
}

// WithJSONEncoder sets a custom JSON marshal function.
func (config *Config) WithJSONEncoder(v func(any) ([]byte, error)) *Config {
	config.JSONEncoder = v
	return config
}

// WithJSONDecoder sets a custom JSON unmarshal function.
func (config *Config) WithJSONDecoder(v func([]byte, any) error) *Config {
	config.JSONDecoder = v
	return config
}

// WithXMLEncoder sets a custom XML marshal function.
func (config *Config) WithXMLEncoder(v func(any) ([]byte, error)) *Config {
	config.XMLEncoder = v
	return config
}

// WithTrustProxy enables trusted proxy header validation.
func (config *Config) WithTrustProxy(v bool) *Config {
	config.TrustProxy = v
	return config
}

// WithTrustedProxies sets the list of trusted proxy IP addresses.
func (config *Config) WithTrustedProxies(v []string) *Config {
	config.TrustProxyConfig.Proxies = v
	return config
}

// WithEnableIPValidation enables validation of IP addresses returned by c.IP().
func (config *Config) WithEnableIPValidation(v bool) *Config {
	config.EnableIPValidation = v
	return config
}

// WithColorScheme sets a custom color scheme for startup messages.
func (config *Config) WithColorScheme(v fiber.Colors) *Config {
	config.ColorScheme = v
	return config
}

// WithRequestMethods sets the list of accepted HTTP methods.
func (config *Config) WithRequestMethods(v []string) *Config {
	config.RequestMethods = v
	return config
}

// WithEnableSplittingOnParsers splits comma-separated query/body/header values.
func (config *Config) WithEnableSplittingOnParsers(v bool) *Config {
	config.EnableSplittingOnParsers = v
	return config
}

// URL returns the server address as "host:port".
func (config *Config) URL() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}

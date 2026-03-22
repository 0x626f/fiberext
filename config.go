package fiberext

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

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

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) WithController(controller *Controller) *Config {
	config.Controllers = append(config.Controllers, controller)
	return config
}

func (config *Config) WithResource(resource *Resource) *Config {
	config.Resources = append(config.Resources, resource)
	return config
}

func (config *Config) WithMiddleware(middleware Handler) *Config {
	config.Middlewares = append(config.Middlewares, middleware)
	return config
}

func (config *Config) WithTLS(v bool) *Config {
	config.TLS = v
	return config
}

func (config *Config) WithMutualTLS(v bool) *Config {
	config.MutualTLS = v
	return config
}

func (config *Config) WithCertFile(v string) *Config {
	config.CertFile = v
	return config
}

func (config *Config) WithKeyFile(v string) *Config {
	config.KeyFile = v
	return config
}

func (config *Config) WithClientCertFile(v string) *Config {
	config.ClientCertFile = v
	return config
}

func (config *Config) WithCertificate(v tls.Certificate) *Config {
	config.Certificate = &v
	return config
}

func (config *Config) WithClientCerts(v *x509.CertPool) *Config {
	config.ClientCerts = v
	return config
}

func (config *Config) WithHost(v string) *Config {
	config.Host = v
	return config
}

func (config *Config) WithPort(v int) *Config {
	config.Port = v
	return config
}

func (config *Config) WithPrefork(v bool) *Config {
	config.Prefork = v
	return config
}

func (config *Config) WithServerHeader(v string) *Config {
	config.ServerHeader = v
	return config
}

func (config *Config) WithStrictRouting(v bool) *Config {
	config.StrictRouting = v
	return config
}

func (config *Config) WithCaseSensitive(v bool) *Config {
	config.CaseSensitive = v
	return config
}

func (config *Config) WithImmutable(v bool) *Config {
	config.Immutable = v
	return config
}

func (config *Config) WithUnescapePath(v bool) *Config {
	config.UnescapePath = v
	return config
}

func (config *Config) WithETag(v bool) *Config {
	config.ETag = v
	return config
}

func (config *Config) WithBodyLimit(v int) *Config {
	config.BodyLimit = v
	return config
}

func (config *Config) WithConcurrency(v int) *Config {
	config.Concurrency = v
	return config
}

func (config *Config) WithViews(v fiber.Views) *Config {
	config.Views = v
	return config
}

func (config *Config) WithViewsLayout(v string) *Config {
	config.ViewsLayout = v
	return config
}

func (config *Config) WithPassLocalsToViews(v bool) *Config {
	config.PassLocalsToViews = v
	return config
}

func (config *Config) WithReadTimeout(v time.Duration) *Config {
	config.ReadTimeout = v
	return config
}

func (config *Config) WithWriteTimeout(v time.Duration) *Config {
	config.WriteTimeout = v
	return config
}

func (config *Config) WithIdleTimeout(v time.Duration) *Config {
	config.IdleTimeout = v
	return config
}

func (config *Config) WithReadBufferSize(v int) *Config {
	config.ReadBufferSize = v
	return config
}

func (config *Config) WithWriteBufferSize(v int) *Config {
	config.WriteBufferSize = v
	return config
}

func (config *Config) WithCompressedFileSuffix(v string) *Config {
	config.CompressedFileSuffix = v
	return config
}

func (config *Config) WithProxyHeader(v string) *Config {
	config.ProxyHeader = v
	return config
}

func (config *Config) WithGETOnly(v bool) *Config {
	config.GETOnly = v
	return config
}

func (config *Config) WithErrorHandler(v ErrorHandler) *Config {
	config.ErrorHandler = v
	return config
}

func (config *Config) WithDisableKeepalive(v bool) *Config {
	config.DisableKeepalive = v
	return config
}

func (config *Config) WithDisableDefaultDate(v bool) *Config {
	config.DisableDefaultDate = v
	return config
}

func (config *Config) WithDisableDefaultContentType(v bool) *Config {
	config.DisableDefaultContentType = v
	return config
}

func (config *Config) WithDisableHeaderNormalizing(v bool) *Config {
	config.DisableHeaderNormalizing = v
	return config
}

func (config *Config) WithDisableStartupMessage(v bool) *Config {
	config.DisableStartupMessage = v
	return config
}

func (config *Config) WithAppName(v string) *Config {
	config.AppName = v
	return config
}

func (config *Config) WithStreamRequestBody(v bool) *Config {
	config.StreamRequestBody = v
	return config
}

func (config *Config) WithDisablePreParseMultipartForm(v bool) *Config {
	config.DisablePreParseMultipartForm = v
	return config
}

func (config *Config) WithReduceMemoryUsage(v bool) *Config {
	config.ReduceMemoryUsage = v
	return config
}

func (config *Config) WithJSONEncoder(v utils.JSONMarshal) *Config {
	config.JSONEncoder = v
	return config
}

func (config *Config) WithJSONDecoder(v utils.JSONUnmarshal) *Config {
	config.JSONDecoder = v
	return config
}

func (config *Config) WithXMLEncoder(v utils.XMLMarshal) *Config {
	config.XMLEncoder = v
	return config
}

func (config *Config) WithNetwork(v string) *Config {
	config.Network = v
	return config
}

func (config *Config) WithEnableTrustedProxyCheck(v bool) *Config {
	config.EnableTrustedProxyCheck = v
	return config
}

func (config *Config) WithTrustedProxies(v []string) *Config {
	config.TrustedProxies = v
	return config
}

func (config *Config) WithEnableIPValidation(v bool) *Config {
	config.EnableIPValidation = v
	return config
}

func (config *Config) WithEnablePrintRoutes(v bool) *Config {
	config.EnablePrintRoutes = v
	return config
}

func (config *Config) WithColorScheme(v fiber.Colors) *Config {
	config.ColorScheme = v
	return config
}

func (config *Config) WithRequestMethods(v []string) *Config {
	config.RequestMethods = v
	return config
}

func (config *Config) WithEnableSplittingOnParsers(v bool) *Config {
	config.EnableSplittingOnParsers = v
	return config
}

func (config *Config) URL() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}

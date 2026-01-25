package fiberext

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	Host string `env:"HOST" default:"localhost"`
	Port int    `env:"PORT" default:"8080"`

	Prefork         bool          `env:"PREFORK" default:"false"`
	BodyLimit       int           `env:"BODY_LIMIT" default:"4194304"`
	Concurrency     int           `env:"CONCURRENCY" default:"262144"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT" default:"0"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT" default:"0"`
	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT" default:"0"`
	ReadBufferSize  int           `env:"READ_BUFFER_SIZE" default:"4096"`
	WriteBufferSize int           `env:"WRITE_BUFFER_SIZE" default:"4096"`

	Bootstrap *BootstrapConfig `env:"-"`
}

func (config *ServerConfig) Convert() fiber.Config {
	fiberConfig := fiber.Config{
		DisableStartupMessage: true,
		Prefork:               config.Prefork,
		BodyLimit:             config.BodyLimit,
		Concurrency:           config.Concurrency,
		ReadTimeout:           config.ReadTimeout,
		WriteTimeout:          config.WriteTimeout,
		IdleTimeout:           config.IdleTimeout,
		ReadBufferSize:        config.ReadBufferSize,
		WriteBufferSize:       config.WriteBufferSize,
	}

	if config.Bootstrap != nil {
		fiberConfig.ErrorHandler = config.Bootstrap.ErrorHandler
	}

	return fiberConfig
}

func (config *ServerConfig) URL() string {
	return fmt.Sprintf("%v:%v", config.Host, config.Port)
}

func (config *ServerConfig) Middlewares() []fiber.Handler {
	if config.Bootstrap == nil {
		return nil
	}

	return config.Bootstrap.Middlewares
}

func (config *ServerConfig) Endpoints() []*EndpointConfig {
	if config.Bootstrap == nil {
		return nil
	}

	return config.Bootstrap.Endpoints
}

func (config *ServerConfig) validateBootstrapConfig() {
	if config.Bootstrap == nil {
		config.Bootstrap = &BootstrapConfig{}
	}
}

func (config *ServerConfig) AddEndpoint(endpoint *EndpointConfig) {
	config.validateBootstrapConfig()
	config.Bootstrap.Endpoints = append(config.Bootstrap.Endpoints, endpoint)
}

func (config *ServerConfig) SetErrorHandler(handler fiber.ErrorHandler) {
	config.validateBootstrapConfig()
	config.Bootstrap.ErrorHandler = handler
}

func (config *ServerConfig) AddMiddleware(handler fiber.Handler) {
	config.validateBootstrapConfig()
	config.Bootstrap.Middlewares = append(config.Bootstrap.Middlewares, handler)
}

type EndpointConfig struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

type BootstrapConfig struct {
	Endpoints    []*EndpointConfig
	ErrorHandler fiber.ErrorHandler
	Middlewares  []fiber.Handler
}

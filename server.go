package fiberext

import (
	"context"
	"crypto/tls"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

// Run builds a fiber server from cfg, registers middlewares, resources and
// controllers, then starts listening in a background goroutine. Cancelling ctx
// triggers graceful shutdown. Returns the underlying *fiber.App.
func Run(ctx context.Context, config *Config) Server {
	if ctx == nil {
		ctx = context.Background()
	}

	server := fiber.New(config.Config)

	for _, middleware := range config.Middlewares {
		server.Use(middleware)
	}

	for _, resource := range config.Resources {
		if resource.Static {
			server.Get(resource.WebPath+"*", static.New(resource.FilePath, resource.StaticConfig))
		} else {
			server.Add([]string{resource.Method}, resource.Path, resource.Handler)
		}
	}

	for _, controller := range config.Controllers {
		group := server.Group(controller.Path)

		for _, resources := range controller.Resources {
			group.Add([]string{resources.Method}, resources.Path, resources.Handler)
		}
	}

	go func() {
		listenCfg := fiber.ListenConfig{
			GracefulContext:       ctx,
			DisableStartupMessage: config.DisableStartupMessage,
		}

		if config.TLS {
			if config.MutualTLS {
				if config.Certificate != nil && config.ClientCerts != nil {
					listenCfg.TLSConfig = &tls.Config{
						Certificates: []tls.Certificate{*config.Certificate},
						ClientAuth:   tls.RequireAndVerifyClientCert,
						ClientCAs:    config.ClientCerts,
					}
				} else {
					listenCfg.CertFile = config.CertFile
					listenCfg.CertKeyFile = config.KeyFile
					listenCfg.CertClientFile = config.ClientCertFile
				}
			} else {
				if config.Certificate != nil {
					listenCfg.TLSConfig = &tls.Config{
						Certificates: []tls.Certificate{*config.Certificate},
					}
				} else {
					listenCfg.CertFile = config.CertFile
					listenCfg.CertKeyFile = config.KeyFile
				}
			}
		}

		if err := server.Listen(config.URL(), listenCfg); err != nil {
			panic(err)
		}
	}()

	return server
}

package fiberext

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

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
			server.Static(resource.WebPath, resource.FilePath, resource.StaticConfig)
		} else {
			server.Add(resource.Method, resource.Path, resource.Handler)
		}
	}

	for _, controller := range config.Controllers {
		group := server.Group(controller.Path)

		for _, resources := range controller.Resources {
			group.Add(resources.Method, resources.Path, resources.Handler)
		}
	}

	go func() {
		var runner func(addr string) error

		if config.TLS {
			if config.MutualTLS {
				if config.Certificate != nil && config.ClientCerts != nil {
					runner = func(address string) error {
						return server.ListenMutualTLSWithCertificate(address, *config.Certificate, config.ClientCerts)
					}
				} else {
					runner = func(address string) error {
						return server.ListenMutualTLS(address, config.CertFile, config.KeyFile, config.ClientCertFile)
					}
				}
			} else {
				if config.Certificate != nil {
					runner = func(address string) error {
						return server.ListenTLSWithCertificate(address, *config.Certificate)
					}
				} else {
					runner = func(address string) error {
						return server.ListenTLS(address, config.CertFile, config.KeyFile)
					}
				}
			}
		} else {
			runner = func(address string) error {
				return server.Listen(address)
			}
		}

		if err := runner(config.URL()); err != nil {
			panic(err)
		}

		<-ctx.Done()

		if err := server.Shutdown(); err != nil {
			panic(err)
		}
	}()

	return server
}

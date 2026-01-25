package fiberext

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func RunHTTPServer(ctx context.Context, config *ServerConfig) {
	server := fiber.New(config.Convert())

	for _, middleware := range config.Middlewares() {
		server.Use(middleware)
	}

	for _, endpoint := range config.Endpoints() {
		server.Add(endpoint.Method, endpoint.Path, endpoint.Handler)
	}

	go func() {
		if err := server.Listen(config.URL()); err != nil {
			panic(err)
		}

		<-ctx.Done()

		if err := server.Shutdown(); err != nil {
			panic(err)
		}
	}()
}

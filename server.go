package fiberext

import (
	"github.com/gofiber/fiber/v2"
)

func RunHTTPServer(config *ServerConfig) error {
	server := fiber.New(config.Convert())

	for _, middleware := range config.Middlewares() {
		server.Use(middleware)
	}

	for _, endpoint := range config.Endpoints() {
		server.Add(endpoint.Method, endpoint.Path, endpoint.Handler)
	}

	return server.Listen(config.URL())
}

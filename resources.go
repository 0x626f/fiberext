package fiberext

import "github.com/gofiber/fiber/v3"

var HealthCheckResource = NewResource(fiber.MethodGet, "/health", func(ctx Context) error {
	return OK(ctx, fiber.Map{
		"status": "ok",
	})
})

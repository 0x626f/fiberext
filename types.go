// Package fiberext provides a thin opinionated wrapper around fiber/v2,
// offering builder-style configuration, typed request helpers, and HTTP status
// response helpers.
package fiberext

import "github.com/gofiber/fiber/v2"

// Server is an alias for *fiber.App.
type Server = *fiber.App

// Handler is an alias for fiber.Handler.
type Handler = fiber.Handler

// ErrorHandler is an alias for fiber.ErrorHandler.
type ErrorHandler = fiber.ErrorHandler

// Context is an alias for *fiber.Ctx.
type Context = *fiber.Ctx

// Controller groups Resources under a common path prefix.
type Controller struct {
	Path      string
	Resources []*Resource
}

// NewController returns a Controller rooted at path.
func NewController(path string) *Controller {
	return &Controller{
		Path: path,
	}
}

// AddResource appends resource to the controller.
func (controller *Controller) AddResource(resource *Resource) *Controller {
	controller.Resources = append(controller.Resources, resource)
	return controller
}

// AddNewResource creates a Resource from method, path and handler and appends it.
func (controller *Controller) AddNewResource(method, path string, handler Handler) *Controller {
	controller.AddResource(NewResource(method, path, handler))
	return controller
}

// Resource is a single route binding.
type Resource struct {
	Method  string
	Path    string
	Handler Handler

	Static       bool
	WebPath      string
	FilePath     string
	StaticConfig fiber.Static
}

// NewResource returns a Resource for the given HTTP method, path and handler.
func NewResource(method, path string, handler Handler) *Resource {
	return &Resource{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

package fiberext

import "github.com/gofiber/fiber/v2"

type Server = *fiber.App

type Handler = fiber.Handler

type ErrorHandler = fiber.ErrorHandler

type Context = *fiber.Ctx
type Controller struct {
	Path      string
	Resources []*Resource
}

func NewController(path string) *Controller {
	return &Controller{
		Path: path,
	}
}

func (controller *Controller) AddResource(resource *Resource) *Controller {
	controller.Resources = append(controller.Resources, resource)
	return controller
}

func (controller *Controller) AddNewResource(method, path string, handler Handler) *Controller {
	controller.AddResource(NewResource(method, path, handler))
	return controller
}

type Resource struct {
	Method  string
	Path    string
	Handler Handler

	Static       bool
	WebPath      string
	FilePath     string
	StaticConfig fiber.Static
}

func NewResource(method, path string, handler Handler) *Resource {
	return &Resource{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

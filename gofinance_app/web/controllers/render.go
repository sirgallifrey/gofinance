package controllers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	// componentHandler := templ.Handler(component)
	// for _, o := range options {
	// 	o(componentHandler)
	// }
	// return adaptor.HTTPHandler(componentHandler)(c)

	c.Set("Content-Type", "text/html")

	return component.Render(c.Context(), c.Response().BodyWriter())
}

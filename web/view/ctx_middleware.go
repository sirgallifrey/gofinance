package view

import "github.com/gofiber/fiber/v2"

func NewViewContext(env string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("viewCtx", &ViewCtx{
			Path: ctx.Path(),
			Env:  env,
		})
		return ctx.Next()
	}
}

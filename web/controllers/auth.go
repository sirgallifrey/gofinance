package controllers

import (
	"requirementor/web/view/pages/auth_pages"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
}

func (c *authController) Index(ctx *fiber.Ctx) error {
	return Render(ctx, auth_pages.SignInPage())

}

func NewAuthController(router fiber.Router) *authController {
	var handler = &authController{}
	router.Get("/", handler.Index)
	return handler
}

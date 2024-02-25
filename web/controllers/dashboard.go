package controllers

import (
	"gofinance/web/view/pages"

	"github.com/gofiber/fiber/v2"
)

type dashboardController struct {
}

func (c *dashboardController) Index(ctx *fiber.Ctx) error {
	return Render(ctx, pages.Test("dashboard"))
}

func NewDashboardController(router fiber.Router) *dashboardController {
	var handler = &dashboardController{}
	router.Get("/", handler.Index)
	return handler
}

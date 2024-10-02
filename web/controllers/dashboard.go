package controllers

import (
	"requirementor/domain"
	"requirementor/web/view/pages"

	"github.com/gofiber/fiber/v2"
)

type dashboardController struct {
}

func (c *dashboardController) Index(ctx *fiber.Ctx) error {
	return Render(ctx, pages.Dashboard("Dashboard"))
}

func (c *dashboardController) Projects(ctx *fiber.Ctx) error {
	return Render(ctx, pages.ProjectsPage{
		Projects: []domain.Project{},
	}.Comp())
}

func NewDashboardController(router fiber.Router) *dashboardController {
	var handler = &dashboardController{}
	router.Get("/", handler.Index)
	router.Get("/projects", handler.Projects)
	return handler
}

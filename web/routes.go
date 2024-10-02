package web

import (
	"requirementor/web/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupControllers(app *fiber.App) {
	controllers.NewDashboardController(app.Group("/"))
}

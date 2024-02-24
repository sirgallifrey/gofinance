package web

import (
	"gofinance/web/view"

	"github.com/gofiber/fiber/v2/log"
)

func SetupWebApp(webApp *webApp) {
	webApp.fiberInstance.Use(view.NewViewContext(webApp.cfg.Env))
	SetupControllers(webApp.fiberInstance)
	if webApp.cfg.IsLocal() {
		log.Info("Setting Up hot-reload websocket")
		WithHotReload(webApp.fiberInstance)
	}
}

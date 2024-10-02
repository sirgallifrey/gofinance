package web

import (
	"requirementor/web/view"

	"github.com/gofiber/fiber/v2/log"
)

func SetupWebApp(webApp *WebApp) {
	webApp.fiberInstance.Use(view.NewViewContext(webApp.Cfg.Env))
	SetupControllers(webApp.fiberInstance)
	if webApp.Cfg.IsLocal() {
		log.Info("Setting Up hot-reload websocket")
		WithHotReload(webApp.fiberInstance)
	}
}

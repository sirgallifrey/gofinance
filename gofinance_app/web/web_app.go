package web

import (
	"fmt"
	"gofinance/config"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type webApp struct {
	cfg           config.AppCfg
	fiberInstance *fiber.App
	logger        zerolog.Logger
}

func (webApp *webApp) Run() {
	webApp.logger.Info().Msg(fmt.Sprintf("Starting server at: %s", webApp.cfg.HTTP.Addr()))
	err := webApp.runHTTP()
	if err != nil {
		webApp.logger.Fatal().Err(err).Msg("Error while trying to run web app.")
	}
}

func (webApp *webApp) runHTTP() error {
	TLS := &webApp.cfg.TLS
	HTTP := &webApp.cfg.HTTP
	if TLS.Cert.Filepath != "" &&
		TLS.Key.Filepath != "" {
		return webApp.fiberInstance.ListenTLS(
			HTTP.Addr(),
			TLS.Cert.Filepath,
			TLS.Key.Filepath,
		)
	} else {
		return webApp.fiberInstance.Listen(HTTP.Addr())
	}
}

func (webApp *webApp) Shutdown() {
	webApp.logger.Info().Msg("Gracefully shutting down...")
	err := webApp.fiberInstance.Shutdown()
	if err != nil {
		webApp.logger.Fatal().Err(err).Msg("Could not shutdown web app.")
	}
}

func setupLogger(cfg *config.AppCfg) zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func Create(cfg *config.AppCfg) webApp {
	app := webApp{
		cfg:           *cfg,
		logger:        setupLogger(cfg),
		fiberInstance: fiber.New(),
	}
	app.logger.Info().Msg(fmt.Sprint("env ", cfg.Env))
	SetupWebApp(&app)
	return app
}

package web

import (
	"context"
	"fmt"
	"os"
	"requirementor/config"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type WebApp struct {
	Cfg           config.AppCfg
	fiberInstance *fiber.App
	logger        zerolog.Logger
	DB            *bun.DB
}

func (webApp *WebApp) Run() {
	webApp.logger.Info().Msg(fmt.Sprintf("Starting server at: %s", webApp.Cfg.HTTP.Addr()))
	err := webApp.runHTTP()
	if err != nil {
		webApp.logger.Fatal().Err(err).Msg("Error while trying to run web app.")
	}
}

func (webApp *WebApp) runHTTP() error {
	TLS := &webApp.Cfg.TLS
	HTTP := &webApp.Cfg.HTTP
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

func (webApp *WebApp) Shutdown() {
	webApp.logger.Info().Msg("Gracefully shutting down...")
	err := webApp.fiberInstance.Shutdown()
	webApp.DB.Close()
	if err != nil {
		webApp.logger.Fatal().Err(err).Msg("Could not shutdown web app.")
	}
}

func setupLogger(cfg *config.AppCfg) zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func setupConn(cfg *config.AppCfg, logger zerolog.Logger) *bun.DB {
	pool, err := pgxpool.New(context.Background(), cfg.Postgres.ConfigString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		panic(err)
	}

	sqldb := stdlib.OpenDBFromPool(pool)
	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}

func Create(cfg *config.AppCfg) WebApp {
	logger := setupLogger(cfg)
	app := WebApp{
		Cfg:           *cfg,
		logger:        setupLogger(cfg),
		fiberInstance: fiber.New(),
		DB:            setupConn(cfg, logger),
	}
	app.logger.Info().Msg(fmt.Sprint("env ", cfg.Env))
	SetupWebApp(&app)
	return app
}

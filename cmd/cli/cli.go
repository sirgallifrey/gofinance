package main

import (
	"fmt"
	"log"
	"os"
	"requirementor/config"
	"requirementor/database/migrations"
	"requirementor/web"
	"strings"

	"github.com/uptrace/bun/migrate"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "requirementor",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "env",
				Value: "dev",
				Usage: "environment",
			},
			&cli.StringFlag{
				Name:      "config",
				TakesFile: true,
			},
		},
		Commands: []*cli.Command{
			dbCommand(migrations.Migrations),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getAppConfig(configFilepath string) *config.AppCfg {
	cfg, err := config.NewAppCfg(configFilepath)
	if err != nil {
		log.Fatalf("cannot load config: %s", err)
		panic("Could not load config. Shutting down")
	}
	return cfg
}

func createApp(configFilepath string) web.WebApp {
	cfg := getAppConfig(configFilepath)
	return web.Create(cfg)
}

func dbCommand(migrations *migrate.Migrations) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "manage database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					app := createApp(c.String("config"))
					defer app.Shutdown()

					migrator := migrate.NewMigrator(app.DB, migrations)
					return migrator.Init(c.Context)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					app := createApp(c.String("config"))
					defer app.Shutdown()

					migrator := migrate.NewMigrator(app.DB, migrations)

					group, err := migrator.Migrate(c.Context)
					if err != nil {
						return err
					}

					if group.ID == 0 {
						fmt.Printf("there are no new migrations to run\n")
						return nil
					}

					fmt.Printf("migrated to %s\n", group)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					app := createApp(c.String("config"))
					defer app.Shutdown()

					migrator := migrate.NewMigrator(app.DB, migrations)

					group, err := migrator.Rollback(c.Context)
					if err != nil {
						return err
					}

					if group.ID == 0 {
						fmt.Printf("there are no groups to roll back\n")
						return nil
					}

					fmt.Printf("rolled back %s\n", group)
					return nil
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					app := createApp(c.String("config"))
					defer app.Shutdown()

					migrator := migrate.NewMigrator(app.DB, migrations)
					return migrator.Lock(c.Context)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					app := createApp(c.String("config"))
					defer app.Shutdown()

					migrator := migrate.NewMigrator(app.DB, migrations)
					return migrator.Unlock(c.Context)
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					app := createApp(c.String("config"))
					defer app.Shutdown()

					migrator := migrate.NewMigrator(app.DB, migrations)

					name := strings.Join(c.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(c.Context, name)
					if err != nil {
						return err
					}
					fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)

					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					app := createApp(c.String("config"))
					defer app.Shutdown()

					migrator := migrate.NewMigrator(app.DB, migrations)

					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					app := createApp(c.String("config"))
					defer app.Shutdown()

					migrator := migrate.NewMigrator(app.DB, migrations)
					ms, err := migrator.MigrationsWithStatus(c.Context)
					if err != nil {
						return err
					}
					fmt.Printf("migrations: %s\n", ms)
					fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
					fmt.Printf("last migration group: %s\n", ms.LastGroup())

					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					app := createApp(c.String("config"))
					defer app.Shutdown()

					migrator := migrate.NewMigrator(app.DB, migrations)

					group, err := migrator.Migrate(c.Context, migrate.WithNopMigration())
					if err != nil {
						return err
					}

					if group.ID == 0 {
						fmt.Printf("there are no new migrations to mark as applied\n")
						return nil
					}

					fmt.Printf("marked as applied %s\n", group)
					return nil
				},
			},
		},
	}
}

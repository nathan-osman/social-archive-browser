package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/social-archive-browser/db"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "social-archive-browser",
		Usage: "viewer for social media archives",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "db-host",
				Value:   "postgres",
				EnvVars: []string{"DB_HOST"},
				Usage:   "PostgreSQL database host",
			},
			&cli.IntFlag{
				Name:    "db-port",
				Value:   5432,
				EnvVars: []string{"DB_PORT"},
				Usage:   "PostgreSQL database port",
			},
			&cli.StringFlag{
				Name:    "db-name",
				Value:   "postgres",
				EnvVars: []string{"DB_NAME"},
				Usage:   "PostgreSQL database name",
			},
			&cli.StringFlag{
				Name:    "db-user",
				Value:   "postgres",
				EnvVars: []string{"DB_USER"},
				Usage:   "PostgreSQL database user",
			},
			&cli.StringFlag{
				Name:    "db-password",
				Value:   "postgres",
				EnvVars: []string{"DB_PASSWORD"},
				Usage:   "PostgreSQL database password",
			},
			&cli.StringFlag{
				Name:    "server-addr",
				Value:   ":http",
				EnvVars: []string{"SERVER_ADDR"},
				Usage:   "address for the server to listen on",
			},
		},
		Action: func(c *cli.Context) error {

			// Create the database connection
			db, err := db.New(
				&db.Config{
					Host:     c.String("db-host"),
					Port:     c.Int("db-port"),
					Name:     c.String("db-name"),
					User:     c.String("db-user"),
					Password: c.String("db-password"),
				},
			)
			if err != nil {
				return err
			}
			defer db.Close()

			// Wait for SIGINT or SIGTERM
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			<-sigChan

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err.Error())
	}
}

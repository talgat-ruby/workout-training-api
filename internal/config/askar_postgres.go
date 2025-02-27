package config

import (
	"context"
	"flag"
)

type AskarPostgresConfig struct {
	DSN string
}

func newAskarPostgresConfig(_ context.Context) *AskarPostgresConfig {
	var c AskarPostgresConfig

	// postgres://workoutdb:admin@localhost/workoutdb?sslmode=disable
	// postgres://postgres:admin@localhost/workoutdb
	// postgres://postgres:admin@localhost/workoutdb?sslmode=disable

	//flag.StringVar(&c.Host, "localhost", c.Host, "postgreSQL host [PG_HOST]")
	//flag.IntVar(&c.Port, "pg-port", c.Port, "postgreSQL port [PG_PORT]")
	//flag.StringVar(&c.Name, "pg-name", c.Name, "postgreSQL name [PG_NAME]")
	//flag.StringVar(&c.User, "pg-user", c.User, "postgreSQL user [PG_USER]")
	//flag.StringVar(&c.Password, "pg-password", c.Password, "postgreSQL password [PG_PASSWORD]")

	flag.StringVar(&c.DSN, "askar-db-dsn", "postgres://postgres:admin@localhost/workoutdb?sslmode=disable", "Askar PostgreSQL DSN")

	flag.Parse()

	return &c
}

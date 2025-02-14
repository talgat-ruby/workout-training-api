package config

import (
	"context"
	"flag"
	"os"
	"strconv"
)

type PostgresConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func newPostgresConfig(_ context.Context) *PostgresConfig {
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))

	// c := &PostgresConfig{
	// 	Host:     os.Getenv("PG_HOST"),
	// 	Port:     port,
	// 	Name:     os.Getenv("PG_NAME"),
	// 	User:     os.Getenv("PG_USER"),
	// 	Password: os.Getenv("PG_PASSWORD"),
	// }

	// flag.StringVar(&c.Host, "pg-host", c.Host, "postgreSQL host [PG_HOST]")
	// flag.IntVar(&c.Port, "pg-port", c.Port, "postgreSQL port [PG_PORT]")
	// flag.StringVar(&c.Name, "pg-name", c.Name, "postgreSQL name [PG_NAME]")
	// flag.StringVar(&c.User, "pg-user", c.User, "postgreSQL user [PG_USER]")
	// flag.StringVar(&c.Password, "pg-password", c.Password, "postgreSQL password [PG_PASSWORD]")

	c := &PostgresConfig{
		Host:     "localhost",
		Port:     port,
		Name:     "postgres",
		User:     "postgres",
		Password: "postgres",
	}

	// postgres://workoutdb:admin@localhost/workoutdb?sslmode=disable
	// postgres://postgres:admin@localhost/workoutdb
	// postgres://postgres:admin@localhost/workoutdb?sslmode=disable

	flag.StringVar(&c.Host, "localhost", c.Host, "postgreSQL host [PG_HOST]")
	flag.IntVar(&c.Port, "pg-port", c.Port, "postgreSQL port [PG_PORT]")
	flag.StringVar(&c.Name, "pg-name", c.Name, "postgreSQL name [PG_NAME]")
	flag.StringVar(&c.User, "pg-user", c.User, "postgreSQL user [PG_USER]")
	flag.StringVar(&c.Password, "pg-password", c.Password, "postgreSQL password [PG_PASSWORD]")

	flag.Parse()

	return c
}

package auth

import (
	"database/sql"
	"log/slog"
	"workout-training-api/internal/config"
)

type Auth struct {
	conf   *config.PostgresConfig
	logger *slog.Logger
	db     *sql.DB
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Auth {
	return &Auth{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}

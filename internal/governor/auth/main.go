package auth

import (
	"log/slog"
	"workout-training-api/internal/config"
	"workout-training-api/internal/types/database"
)

type Auth struct {
	conf   *config.Config
	logger *slog.Logger
	db     database.Database
}

func New(conf *config.Config, logger *slog.Logger, db database.Database) *Auth {
	return &Auth{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}

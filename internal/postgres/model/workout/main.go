package workout

import (
	"database/sql"
	"log/slog"
	"workout-training-api/internal/config"
)

type Workout struct {
	conf   *config.PostgresConfig
	logger *slog.Logger
	db     *sql.DB
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Workout {
	return &Workout{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}

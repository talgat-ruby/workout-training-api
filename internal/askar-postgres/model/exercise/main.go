package exercise

import (
	"database/sql"
	"log/slog"
	"workout-training-api/internal/config"
)

type ExerciseModel struct {
	conf   *config.AskarPostgresConfig
	logger *slog.Logger
	db     *sql.DB
}

func New(conf *config.AskarPostgresConfig, logger *slog.Logger, db *sql.DB) *ExerciseModel {
	return &ExerciseModel{conf: conf, logger: logger, db: db}
}

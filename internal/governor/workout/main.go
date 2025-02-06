package workout

import (
	"log/slog"
	"workout-training-api/internal/config"
	"workout-training-api/internal/types/database"
)

type Workout struct {
	conf   *config.Config
	logger *slog.Logger
	db     database.Database
}

func New(conf *config.Config, logger *slog.Logger, db database.Database) *Workout {
	return &Workout{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}

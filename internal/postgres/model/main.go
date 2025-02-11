package model

import (
	"database/sql"
	"log/slog"
	"workout-training-api/internal/config"
	"workout-training-api/internal/postgres/model/auth"
	"workout-training-api/internal/postgres/model/workout"
)

type Model struct {
	*auth.Auth
	*workout.Workout
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Model {
	return &Model{
		Auth:    auth.New(conf, logger.With(slog.String("component", "auth")), db),
		Workout: workout.New(conf, logger.With(slog.String("component", "workout")), db),
	}
}

package model

import (
	"database/sql"
	"log/slog"
	"workout-training-api/internal/askar-postgres/model/exercise"
	"workout-training-api/internal/config"
)

type Model struct {
	*exercise.ExerciseModel
}

func New(conf *config.PostgresConfig, logger *slog.Logger, db *sql.DB) *Model {

	return &Model{ExerciseModel: exercise.New(conf, logger.With(slog.String("component", "exercise")), db)}
}

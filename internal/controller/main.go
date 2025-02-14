package controller

import (
	"workout-training-api/internal/config"
	"workout-training-api/internal/controller/auth"
	"workout-training-api/internal/controller/workout"
	"workout-training-api/internal/types/database"
)

type Controller struct {
	*auth.AuthController
	*workout.WorkoutController
}

func New(conf *config.Config, db database.Database) *Controller {
	return &Controller{
		WorkoutController: workout.New(db),
	}
}

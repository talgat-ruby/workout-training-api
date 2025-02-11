package workout

import (
	"workout-training-api/internal/types/database"
)

type WorkoutController struct {
	db database.Workout
}

func New() *WorkoutController {
	return &WorkoutController{}
}

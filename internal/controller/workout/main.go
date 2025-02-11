package workout

import (
	"workout-training-api/internal/types/database"
)

type WorkoutController struct {
	db database.Database
}

func New(db database.Database) *WorkoutController {
	return &WorkoutController{
		db: db,
	}
}

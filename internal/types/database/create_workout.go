package database

import (
	"time"
	"workout-training-api/internal/postgres/db_types/workout"
)

type CreateWorkoutReq interface {
	GetUserID() string
	GetName() string
	GetExercises() []Exercise
	GetDate() time.Time
	GetExercise() []string
	GetScheduledTimes() []time.Time
}

type CreateWorkoutResp interface {
	GetWorkout() workout.Workout
}

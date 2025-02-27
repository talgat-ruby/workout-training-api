package database

import (
	"time"
	"workout-training-api/internal/constant"
	"workout-training-api/internal/postgres/db_types/workout"
)

type CreateWorkoutReq interface {
	GetUserID() string
	GetName() string
	GetExercises() []workout.Exercise
	GetStatus() constant.WorkoutStatus
	GetDescription() string
	GetScheduledDate() []time.Time
}

type CreateWorkoutResp interface {
	GetWorkout() workout.Workout
}

package controller

import (
	"time"
	"workout-training-api/internal/postgres/db_types/workout"
)

type CreateWorkoutReq interface {
	GetName() string
	GetDescription() string
	GetExercises() []workout.Exercise
	GetDate() time.Time
	GetScheduledDate() []time.Time
}

type CreateWorkoutResp interface {
	GetWorkout() workout.Workout
}

type Exercise interface {
	GetName() string
	GetDescription() string
	GetMuscleGroup() string
	GetCategory() string
	GetRepetitions() int
	GetSets() int
	GetWeight() float64
}

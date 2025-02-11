package controller

import (
	"time"

	"workout-training-api/internal/constant"
	"workout-training-api/internal/postgres/db_types/workout"
)

type CreateWorkoutReq interface {
	GetName() string
	GetDescription() string
	GetExercises() []Exercise
	GetStatus() constant.WorkoutStatus
	GetScheduledDate() []time.Time
}

type CreateWorkoutResp interface {
	GetWorkout() workout.Workout
}

type Exercise interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetMuscleGroup() string
	GetCategory() string
	GetRepetitions() int
	GetSets() int
	GetWeight() float64
}

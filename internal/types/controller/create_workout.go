package controller

import (
	"time"

	"workout-training-api/internal/constant"
)

type CreateWorkoutReq interface {
	GetName() string
	GetDescription() string
	GetExercises() []Exercise
	GetStatus() constant.WorkoutStatus
	ScheduledDate() time.Time
}

type CreateWorkoutResp interface {
	GetID() string
	GetStatus() constant.WorkoutStatus
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

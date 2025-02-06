package controller

import "time"

type CreateWorkoutReq interface {
	GetName() string
	GetDescription() string
	GetExercises() []Exercise
	GetDate() time.Time
}

type CreateWorkoutResp interface {
	GetWorkout() Workout
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

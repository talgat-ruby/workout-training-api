package controller

import "time"

type CreateWorkoutReq interface {
	GetNameWorkout() string
	GetDescription() string
	GetExercises() []Exercise
	GetDate() []time.Time
}

type CreateWorkoutResp interface{}

type Exercise interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetCategory() string
	GetRepetitions() int
	GetSets() int
	GetWeight() float64
}
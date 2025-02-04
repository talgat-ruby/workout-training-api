package controller

import "time"

type CreateWorkoutReq interface {
	GetExercises() []Exercise
	GetDate() time.Time
}

type CreateWorkoutResp interface{}

type Exercise interface {
	GetName() string
	GetDescription() string
	GetCategory() string
	GetRepetitions() int
	GetSets() int
	GetWeight() float64
}

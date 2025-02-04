package controller

import (
	"time"
)

type ListWorkoutsReq interface {
	GetUserID() string
	GetStatus() string
}

type ListWorkoutsResp interface {
	GetWorkouts() []Workout
}

type Workout interface {
	GetID() string
	GetExercises() []Exercise
	GetDate() time.Time
	GetComments() string
}

package controller

import "time"

type ListWorkoutsReq interface {
	GetUserID() string
}

type ListWorkoutsResp interface {
	GetWorkouts() []Workout
}

type Workouts interface {
	GetID() string
	GetExercises() []Exercise
	GetDate() time.Time
	GetComments() string
}

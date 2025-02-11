package controller

import (
	"workout-training-api/internal/postgres/db_types/workout"
)

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

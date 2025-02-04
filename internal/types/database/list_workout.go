package database

import "time"

type ListWorkoutReq interface {
	GetUserID() string
}

type ListWorkoutResp interface {
	GetList() []WorkoutListResp
}

type WorkoutListResp interface {
	GetID() string
	GetUserID() string
	GetStatus() string
	GetScheduledTimes() []time.Time
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetExercises() []ExerciseResp
}

package database

import "time"

type ListWorkoutReq interface {
	GetUserID() string
}

type ListWorkoutResp interface {
	GetWorkouts() []WorkoutResp
}

type WorkoutResp interface {
	GetID() string
	GetUserID() string
	GetStatus() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

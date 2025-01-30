package database

import "time"

type CreateWorkoutReq interface {
	GetUserID() string
	GetExerciseIDs() []string
	GetScheduledTimes() []time.Time
}

type CreateWorkoutResp interface {
}

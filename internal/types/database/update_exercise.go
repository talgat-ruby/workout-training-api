package database

import "time"

type UpdateExerciseReq interface {
	GetName() string
	GetDescription() string
	GetCategories() []string
	GetMuscleGroups() []string
	GetScheduledTimes() []time.Time
}

type UpdateExerciseResp interface {
}

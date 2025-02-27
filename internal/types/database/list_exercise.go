package database

import "time"

type ListExerciseReq interface {
}

type ListExerciseResp interface {
	GetList() []ExerciseResp
}

type ExerciseResp interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetCategories() []string
	GetMuscleGroups() []string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

package database

type CreateExerciseReq interface {
	GetWorkoutID() string
	GetName() string
	GetDescription() string
	GetRepetitions() int32
	GetSets() int32
	GetWeight() int32
	GetMuscleGroups() []string
	GetCategories() []string
}

type CreateExerciseResp interface{}

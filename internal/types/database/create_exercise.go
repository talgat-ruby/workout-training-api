package database

type CreateExerciseReq interface {
	GetName() string
	GetDescription() string
	GetCategories() []string
	GetMuscleGroups() []string
}

type CreateExerciseResp interface {
}

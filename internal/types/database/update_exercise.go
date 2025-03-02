package database

type UpdateExerciseReq interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetCategories() []string
	GetMuscleGroups() []string
	//GetScheduledTimes() []time.Time
}

type UpdateExerciseResp interface {
}

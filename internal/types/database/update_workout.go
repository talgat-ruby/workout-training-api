package database

type UpdateWorkoutReq interface {
	GetUserID() string
	GetID() string
	GetComments() []string
	GetExerciseIDs() []string
}

type UpdateWorkoutResp interface {
}

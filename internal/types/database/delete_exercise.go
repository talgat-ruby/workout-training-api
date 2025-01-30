package database

type DeleteExerciseReq interface {
	GetUserID() string
	GetID() string
}

type DeleteExerciseResp interface {
}

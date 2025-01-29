package database

type CreateWorkoutReq interface {
	GetUserID() string
	GetExercises() [][]interface{}
}

type CreateWorkoutResp interface {
}

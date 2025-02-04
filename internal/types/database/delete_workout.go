package database

type DeleteWorkoutReq interface {
	GetUserID() string
	GetID() string
}

type DeleteWorkoutResp interface {
}

package database

type UpdateWorkoutReq interface {
	GetUserID() string
	GetID() string
	GetComments() []string
}

type UpdateWorkoutResp interface {
}

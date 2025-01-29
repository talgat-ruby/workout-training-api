package database

type CreateUserReq interface {
	GetEmail() string
	GetPasswordHash() string
}

type CreateUserResp interface {
	GetID() string
}

package database

type FindUserReq interface {
	GetEmail() string
	GetPasswordHash() string
}

type FindUserResp interface {
	GetID() string
	GetEmail() string
}

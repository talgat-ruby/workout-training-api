package database

import "time"

type FindUserReq interface {
	GetEmail() string
}

type FindUserResp interface {
	GetID() string
	GetEmail() string
	GetPasswordHash() string
	GetSalt() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

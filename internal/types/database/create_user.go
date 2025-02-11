package database

import (
	"context"
	"fmt"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash string
}

type CreateUserReq interface {
	GetEmail() string
	GetPasswordHash() string
}

type CreateUserResp interface {
	GetID() string
}

type createUserResp struct {
	ID string
}

func (r createUserResp) GetID() string {
	return r.ID
}

func (d *Database) CreateUser(ctx context.Context, req CreateUserReq) (CreateUserResp, error) {
	user := User{Email: req.GetEmail(), PasswordHash: req.GetPasswordHash()}
	if err := d.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return createUserResp{ID: fmt.Sprintf("%d", user.ID)}, nil
}

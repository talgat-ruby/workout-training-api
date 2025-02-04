package database

import "context"

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

func CreateUser(ctx context.Context, req CreateUserReq) (*User, error) {
	user := User{Email: req.GetEmail(), PasswordHash: req.GetPasswordHash()}
	if err := DB.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

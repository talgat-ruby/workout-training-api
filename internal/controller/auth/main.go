package auth

import (
	"workout-training-api/internal/types/database"
)

type AuthController struct {
	db database.Database
}

func New(db database.Database) *AuthController {
	return &AuthController{
		db: db,
	}
}

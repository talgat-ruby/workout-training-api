package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"workout-training-api/internal/postgres/db_types/workout"
	"workout-training-api/internal/types/database"
)

// FindUserByEmail finds a user by email using a raw SQL query.
func (m *Auth) FindUserByEmail(ctx context.Context, req database.FindUserReq) (database.FindUserResp, error) {
	log := m.logger.With(slog.String("handler", "FindUserByEmail"))

	if req == nil {
		log.ErrorContext(ctx, "request is nil")
		return nil, fmt.Errorf("request is nil")
	}

	// Define the SQL query.
	query := `
		SELECT user_id, email, hashed_password, salt, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	// Prepare a variable to hold the user.
	var user workout.User

	// Execute the query.
	err := m.db.QueryRowContext(ctx, query, req.GetEmail()).Scan(
		&user.UserID,
		&user.Email,
		&user.HashedPassword,
		&user.Salt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.WarnContext(ctx, "no rows found")
			return nil, nil
		}
		log.ErrorContext(ctx, "failed to find user by email", slog.Any("error", err))
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	log.InfoContext(ctx, "success", slog.String("email", user.Email))
	return &findUserByEmailResp{user: user}, nil
}

// findUserByEmailResp wraps the workout.User and implements the database.FindUserResp interface.
type findUserByEmailResp struct {
	user workout.User
}

func (r *findUserByEmailResp) GetID() string {
	return r.user.UserID
}

func (r *findUserByEmailResp) GetEmail() string {
	return r.user.Email
}

func (r *findUserByEmailResp) GetPasswordHash() string {
	return r.user.HashedPassword
}

func (r *findUserByEmailResp) GetSalt() string {
	return r.user.Salt
}

func (r *findUserByEmailResp) GetCreatedAt() time.Time {
	return r.user.CreatedAt
}

func (r *findUserByEmailResp) GetUpdatedAt() time.Time {
	return r.user.UpdatedAt
}

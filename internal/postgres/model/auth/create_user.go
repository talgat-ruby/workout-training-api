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

func (m *Auth) CreateUser(ctx context.Context, req database.CreateUserReq) (database.CreateUserResp, error) {
	log := m.logger.With(slog.String("handler", "CreateUser"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	newUser := workout.User{
		Email:          req.GetEmail(),
		HashedPassword: req.GetPasswordHash(),
		Salt:           req.GetSalt(),
	}
	query := `
        INSERT INTO users (
            email,
            hashed_password,
            salt,
            created_at,
            updated_at
        ) VALUES (
            $1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
        )
        RETURNING 
            user_id,
            email,
            hashed_password,
            salt,
            created_at,
            updated_at`

	err := m.db.QueryRowContext(
		ctx,
		query,
		newUser.Email,
		newUser.HashedPassword,
		newUser.Salt,
	).Scan(
		&newUser.UserID,
		&newUser.Email,
		&newUser.HashedPassword,
		&newUser.Salt,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.WarnContext(ctx, "no rows found")
			return nil, nil
		}
		log.ErrorContext(ctx, "failed to insert user", slog.Any("error", err))
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	log.InfoContext(
		ctx,
		"success",
		slog.String("email", newUser.Email),
	)
	return &createUserResp{newUser}, nil
}

type createUserResp struct {
	workout.User
}

func (resp *createUserResp) GetID() string {
	return fmt.Sprint(resp.UserID)
}

func (resp *createUserResp) GetEmail() string {
	return resp.Email
}

func (resp *createUserResp) GetCreatedAt() time.Time {
	return resp.CreatedAt
}

func (resp *createUserResp) GetUpdatedAt() time.Time {
	return resp.UpdatedAt
}

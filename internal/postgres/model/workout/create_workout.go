package workout

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"log/slog"
	"workout-training-api/internal/postgres/db_types/workout"

	"workout-training-api/internal/types/database"
)

// CreateExpense creates a new expense record using a raw SQL query.
func (m *Workout) CreateWorkout(ctx context.Context, req database.CreateWorkoutReq) (database.CreateWorkoutResp, error) {
	log := m.logger.With(slog.String("handler", "CreateWorkout"))

	// Validate the request.
	if req == nil {
		log.ErrorContext(ctx, "request is nil")
		return nil, fmt.Errorf("request is nil")
	}

	/*
		type CreateWorkoutReq interface {
			GetUserID() string
			GetName() string
			GetExercises() []workout.Exercise
			GetStatus() constant.WorkoutStatus
			GetDescription() string
			GetScheduledDate() []time.Time
		}

	*/
	// In this example, we assume UserID is a string; if you need conversion,
	// do it here. For instance, if your database expects an integer, convert accordingly.
	userID := req.GetUserID()
	var exerciseToSave workout.Exercise

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
	// Need to create workout with slice of exercises inside how to do it in sql:

	queryForExercise := `
	INSERT INTO exercise (workout_id, name, category, muscle_group,sets, reps_per_set, weight, notes) 
	VALUES ($1,$2,$3,$4,$5)
	RETURNING workout_id, name, category, muscle_group,sets, reps_per_set, weight, notes
`
	err := m.db.QueryRowContext(
		ctx,
		queryForExercise,
		exerciseToSave.WorkoutID,
		exerciseToSave.Name,
		exerciseToSave.Category,
		exerciseToSave.MuscleGroup,
		exerciseToSave.RepsPerSet,
		exerciseToSave.WeightKg,
		exerciseToSave.Notes,
	).Scan(
		&exerciseToSave.WorkoutID,
		&exerciseToSave.Name,
		&exerciseToSave.Category,
		&exerciseToSave.MuscleGroup,
		&exerciseToSave.RepsPerSet,
		&exerciseToSave.WeightKg,
		&exerciseToSave.Notes,
		&exerciseToSave.CreatedAt,
		&exerciseToSave.UpdatedAt,
	)
	// Prepare the raw SQL query for inserting a workout.
	// We assume the workouts table has columns: user_id, name, description, status, scheduled_date.
	// The RETURNING clause returns all necessary columns.
	query := `
		INSERT INTO workouts (user_id, name,exercises,description, status, scheduled_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING workout_id, user_id, name, description, status, scheduled_date, created_at, updated_at
	`

	// Define a variable to hold the inserted workout.
	var wk workout.Workout

	// Execute the query and scan the result.
	// For the scheduled_date, we use pq.Array both for input and output.
	err := m.db.QueryRowContext(
		ctx,
		query,
		userID,
		req.GetName(),
		req.GetDescription(),
		req.GetStatus(),
		pq.Array(req.GetScheduledDate()),
	).Scan(
		&wk.WorkoutID,
		&wk.UserID,
		&wk.Name,
		&wk.Description,
		&wk.Status,
		pq.Array(&wk.ScheduledDate),
		&wk.CreatedAt,
		&wk.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.WarnContext(ctx, "no rows returned")
			return nil, nil
		}
		log.ErrorContext(ctx, "failed to insert workout", slog.Any("error", err))
		return nil, fmt.Errorf("failed to insert workout: %w", err)
	}

	log.InfoContext(ctx, "workout inserted successfully", slog.String("workout_id", wk.WorkoutID))
	return &createWorkoutResp{workout: wk}, nil
}

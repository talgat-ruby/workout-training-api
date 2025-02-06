package workout

import (
	"context"
	"database/sql"

	"fmt"
	"log/slog"

	"workout-training-api/internal/constant"
	"workout-training-api/internal/postgres/db_types/workout"
	"workout-training-api/internal/types/database"
)

func (r *Workout) CreateWorkout(ctx context.Context, req database.CreateWorkoutReq) (database.CreateWorkoutResp, error) {
	log := r.logger.With(slog.String("handler", "CreateWorkout"))

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert workout
	var workoutID int
	err = tx.QueryRowContext(ctx, `
		INSERT INTO workouts (user_id, name, status, scheduled_date)
		VALUES ($1, $2, $3, $4)
		RETURNING workout_id
	`, req.GetUserID(), req., constant.StatusPending, req.GetExercise()).Scan(&workoutID)

	if err != nil {
		log.ErrorContext(ctx, "failed to insert workout", slog.Any("error", err))
		return nil, fmt.Errorf("failed to insert workout: %w", err)
	}

	// Insert exercises
	for _, exercise := range req.GetExercise() {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO exercises (
				workout_id,
				muscle_group,
				category,
				name,
				sets,
				reps_per_set,
				weight_kg,
				notes
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`,
			workoutID,
			constant.MuscleGroupChest, // This should be mapped from exercise.GetCategory()
			exercise.GetCategory(),
			exercise.GetName(),
			exercise.GetSets(),
			exercise.GetRepetitions(),
			sql.NullFloat64{Float64: exercise.GetWeight(), Valid: exercise.GetWeight() > 0},
			sql.NullString{String: exercise.GetDescription(), Valid: exercise.GetDescription() != ""},
		)

		if err != nil {
			log.ErrorContext(ctx, "failed to insert exercise", slog.Any("error", err))
			return nil, fmt.Errorf("failed to insert exercise: %w", err)
		}
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		log.ErrorContext(ctx, "failed to commit transaction", slog.Any("error", err))
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Fetch the created workout with exercises
	var result workout.Workout
	err = r.db.QueryRowContext(ctx, `
		SELECT 
			w.workout_id,
			w.user_id,
			w.name,
			w.description,
			w.status,
			w.scheduled_date,
			w.created_at,
			w.updated_at
		FROM workouts w
		WHERE w.workout_id = $1
	`, workoutID).Scan(
		&result.WorkoutID,
		&result.UserID,
		&result.Name,
		&result.Description,
		&result.Status,
		&result.ScheduledDate,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.ErrorContext(ctx, "failed to fetch created workout", slog.Any("error", err))
		return nil, fmt.Errorf("failed to fetch created workout: %w", err)
	}

	// Fetch exercises
	rows, err := r.db.QueryContext(ctx, `
		SELECT 
			exercise_id,
			workout_id,
			muscle_group,
			category,
			name,
			sets,
			reps_per_set,
			weight_kg,
			notes,
			created_at,
			updated_at
		FROM exercises
		WHERE workout_id = $1
	`, workoutID)

	if err != nil {
		log.ErrorContext(ctx, "failed to fetch exercises", slog.Any("error", err))
		return nil, fmt.Errorf("failed to fetch exercises: %w", err)
	}
	defer rows.Close()

	var exercises []workout.Exercise
	for rows.Next() {
		var exercise workout.Exercise
		err := rows.Scan(
			&exercise.ExerciseID,
			&exercise.WorkoutID,
			&exercise.MuscleGroup,
			&exercise.Category,
			&exercise.Name,
			&exercise.Sets,
			&exercise.RepsPerSet,
			&exercise.WeightKg,
			&exercise.Notes,
			&exercise.CreatedAt,
			&exercise.UpdatedAt,
		)
		if err != nil {
			log.ErrorContext(ctx, "failed to scan exercise", slog.Any("error", err))
			return nil, fmt.Errorf("failed to scan exercise: %w", err)
		}
		exercises = append(exercises, exercise)
	}

	result.Exercises = exercises

	return &workoutResponse{
		workout: result,
	}, nil
}

type workoutResponse struct {
	workout workout.Workout
}

func (r *workoutResponse) GetWorkout() workout.Workout {
	return r.workout
}

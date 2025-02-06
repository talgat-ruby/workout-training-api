package workout

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"workout-training-api/internal/authentication"
	"workout-training-api/internal/constant"
	"workout-training-api/internal/postgres/db_types/workout"
	"workout-training-api/internal/types/controller"
	"workout-training-api/internal/types/database"
)

// CreateWorkout implements the controller interface
func (w *Workout) CreateWorkout(ctx context.Context, req controller.CreateWorkoutReq) (*controller.CreateWorkoutResp, error) {
	log := w.logger.With(slog.String("handler", "CreateWorkout"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	// Get user from context
	user, ok := ctx.Value(constant.ContextUser).(*authentication.UserData)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}

	// Create database request
	dbReq := newCreateWorkoutDBReq(user.ID, req)
	dbResp, err := w.db.CreateWorkout(ctx, dbReq)
	if err != nil {
		log.ErrorContext(ctx, "db request failed", slog.Any("error", err))
		return nil, fmt.Errorf("db request failed: %w", err)
	}
	if dbResp == nil {
		return nil, nil
	}

	log.InfoContext(ctx, "success")

	return dbResp.GetWorkout(), nil
}

// Request structs
type createWorkoutDBReq struct {
	userID        string
	name          string
	description   string
	scheduledDate time.Time
	exercises     []createExerciseDBReq
}

type createExerciseDBReq struct {
	name        string
	description string
	muscleGroup constant.MuscleGroup
	category    constant.ExerciseCategory
	repetitions int
	sets        int
	weight      float64
}

// Implement database.CreateWorkoutReq interface
func (req *createWorkoutDBReq) GetUserID() string {
	return req.userID
}

func (req *createWorkoutDBReq) GetName() string {
	return req.name
}

func (req *createWorkoutDBReq) GetDescription() string {
	return req.description
}

func (req *createWorkoutDBReq) GetScheduledDate() time.Time {
	return req.scheduledDate
}

func (req *createWorkoutDBReq) GetExercises() []database.Exercise {
	exercises := make([]database.Exercise, len(req.exercises))
	for i, ex := range req.exercises {
		exercises[i] = &ex
	}
	return exercises
}

// Implement database.Exercise interface
func (ex *createExerciseDBReq) GetName() string {
	return ex.name
}

func (ex *createExerciseDBReq) GetDescription() string {
	return ex.description
}

func (ex *createExerciseDBReq) GetMuscleGroup() constant.MuscleGroup {
	return ex.muscleGroup
}

func (ex *createExerciseDBReq) GetCategory() constant.ExerciseCategory {
	return ex.category
}

func (ex *createExerciseDBReq) GetRepetitions() int {
	return ex.repetitions
}

func (ex *createExerciseDBReq) GetSets() int {
	return ex.sets
}

func (ex *createExerciseDBReq) GetWeight() float64 {
	return ex.weight
}

func newCreateWorkoutDBReq(userID string, req controller.CreateWorkoutReq) database.CreateWorkoutReq {
	dbExercises := make([]createExerciseDBReq, len(req.GetExercises()))
	for i, ex := range req.GetExercises() {
		dbExercises[i] = createExerciseDBReq{
			name:        ex.GetName(),
			description: ex.GetDescription(),
			muscleGroup: constant.MuscleGroup(ex.GetMuscleGroup()),
			category:    constant.ExerciseCategory(ex.GetCategory()),
			repetitions: ex.GetRepetitions(),
			sets:        ex.GetSets(),
			weight:      ex.GetWeight(),
		}
	}

	return &createWorkoutDBReq{
		userID:        userID,
		name:          req.GetName(),
		description:   req.GetDescription(),
		scheduledDate: req.GetScheduledDate(),
		exercises:     dbExercises,
	}
}

// Response struct
type workoutResponse struct {
	workout workout.Workout
}

func (r *workoutResponse) GetWorkout() workout.Workout {
	return r.workout
}

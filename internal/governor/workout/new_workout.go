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
)

// CreateWorkout creates a new workout for the authenticated user.
//
//goland:noinspection ALL
func (w *Workout) CreateWorkout(ctx context.Context, req controller.CreateWorkoutReq) (controller.CreateWorkoutResp, error) {
	log := w.logger.With(slog.String("handler", "CreateWorkout"))

	// Validate the incoming request.
	if err := validateCreateWorkoutRequest(req); err != nil {
		log.ErrorContext(ctx, "invalid request", slog.Any("error", err))
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Get the authenticated user.
	user, ok := ctx.Value(constant.ContextUser).(*authentication.UserData)
	if !ok {
		return nil, fmt.Errorf("failed to get user from context: user authentication required")
	}

	dbReq := newCreateWorkoutDBReq(user.ID, req)
	dbResp, err := w.db.CreateWorkout(ctx, dbReq)
	if err != nil {
		log.ErrorContext(ctx, "db request failed", slog.Any("error", err))
		return nil, fmt.Errorf("failed to create workout: %w", err)
	}
	if dbResp == nil {
		return nil, fmt.Errorf("workout creation failed: no response from database")
	}

	workoutResult := dbResp.GetWorkout()
	log.InfoContext(ctx, "workout created successfully",
		slog.String("workout_id", workoutResult.WorkoutID),
		slog.String("user_id", user.ID),
		slog.Int("exercise_count", len(req.GetExercises())))

	// Wrap the database workout in a type that implements controller.CreateWorkoutResp.
	return &createWorkoutResp{workout: workoutResult}, nil
}

// validateCreateWorkoutRequest ensures that the request has the required fields.
func validateCreateWorkoutRequest(req controller.CreateWorkoutReq) error {
	if req == nil {
		return fmt.Errorf("request is nil")
	}
	if len(req.GetName()) == 0 {
		return fmt.Errorf("workout name is required")
	}
	if len(req.GetExercises()) == 0 {
		return fmt.Errorf("at least one exercise is required")
	}
	if len(req.GetScheduledDate()) == 0 {
		return fmt.Errorf("scheduled date is required")
	}
	return nil
}

// ----------------------------------------------------------------------------
// Implementation of the database request to create a workout.
// This type is expected to implement the database.CreateWorkoutReq interface.
// ----------------------------------------------------------------------------

type createWorkoutDBReq struct {
	userID        string
	name          string
	description   string
	workoutStatus constant.WorkoutStatus
	scheduledDate []time.Time
	exercises     []createExerciseDBReq
}

// GetUserID returns the user ID.
func (r *createWorkoutDBReq) GetUserID() string {
	return r.userID
}

// GetName returns the workout name.
func (r *createWorkoutDBReq) GetName() string {
	return r.name
}

// GetDescription returns the workout description.
func (r *createWorkoutDBReq) GetDescription() string {
	return r.description
}

// GetStatus returns the workout status.
func (r *createWorkoutDBReq) GetStatus() constant.WorkoutStatus {
	return r.workoutStatus
}

// GetScheduledDate returns the full slice of scheduled dates.
// (If scheduledDate is nil, an empty slice is returned.)
func (r *createWorkoutDBReq) GetScheduledDate() []time.Time {
	if r.scheduledDate == nil {
		return []time.Time{}
	}
	return r.scheduledDate
}

// GetExercises converts the internal exercise requests into the database model.
func (r *createWorkoutDBReq) GetExercises() []workout.Exercise {
	exercises := make([]workout.Exercise, len(r.exercises))
	for i, ex := range r.exercises {
		exercises[i] = workout.Exercise{
			Name:        ex.name,
			MuscleGroup: ex.muscleGroup,
			Category:    ex.category,
			Sets:        ex.sets,
			RepsPerSet:  ex.repetitions,
			WeightKg:    ex.weight,
			Notes:       ex.notes,
		}
	}
	return exercises
}

type createExerciseDBReq struct {
	name        string
	muscleGroup constant.MuscleGroup
	category    constant.ExerciseCategory
	repetitions int
	sets        int
	weight      float64
	notes       string
}

func newCreateWorkoutDBReq(userID string, req controller.CreateWorkoutReq) *createWorkoutDBReq {
	exercisesReq := req.GetExercises()
	exercises := make([]createExerciseDBReq, len(exercisesReq))
	for i, ex := range exercisesReq {
		exercises[i] = createExerciseDBReq{
			name:        ex.Name,
			muscleGroup: ex.MuscleGroup,
			category:    ex.Category,
			repetitions: ex.RepsPerSet,
			sets:        ex.Sets,
			weight:      ex.WeightKg,
			notes:       ex.Notes,
		}
	}

	return &createWorkoutDBReq{
		userID:        userID,
		name:          req.GetName(),
		workoutStatus: req.GetStatus(),
		description:   req.GetDescription(),
		scheduledDate: req.GetScheduledDate(),
		exercises:     exercises,
	}
}

type createWorkoutResp struct {
	workout workout.Workout
}

// GetWorkout returns the underlying workout.
func (r *createWorkoutResp) GetWorkout() workout.Workout {
	return r.workout
}

func (r *createWorkoutDBReq) GetDate() []time.Time {
	return r.scheduledDate
}

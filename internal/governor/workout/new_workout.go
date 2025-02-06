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

func (w *Workout) CreateWorkout(ctx context.Context, req controller.CreateWorkoutReq) (*workout.Workout, error) {
	log := w.logger.With(slog.String("handler", "CreateWorkout"))

	if err := validateCreateWorkoutRequest(req); err != nil {
		log.ErrorContext(ctx, "invalid request", slog.Any("error", err))
		return nil, fmt.Errorf("invalid request: %w", err)
	}

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

	return &workoutResult, nil
}

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

type createWorkoutDBReq struct {
	userID        string
	name          string
	description   string
	scheduledDate []time.Time
	exercises     []createExerciseDBReq
}

func (req *createWorkoutDBReq) GetDate() time.Time {
	if len(req.scheduledDate) == 0 {
		return time.Time{}
	}
	return req.scheduledDate[0]
}

func (req *createWorkoutDBReq) GetExercise() []string {
	exercises := make([]string, len(req.exercises))
	for i, ex := range req.exercises {
		exercises[i] = ex.name
	}
	return exercises
}

func (req *createWorkoutDBReq) GetScheduledTimes() []time.Time {
	return req.scheduledDate
}

func (req *createWorkoutDBReq) GetScheduledDate() time.Time {
	if len(req.scheduledDate) == 0 {
		return time.Time{}
	}
	return req.scheduledDate[0]
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

func (req *createWorkoutDBReq) GetUserID() string {
	return req.userID
}

func (req *createWorkoutDBReq) GetName() string {
	return req.name
}

func (req *createWorkoutDBReq) GetDescription() string {
	return req.description
}

func (req *createWorkoutDBReq) GetExercises() []workout.Exercise {
	exercises := make([]workout.Exercise, len(req.exercises))
	for i, ex := range req.exercises {
		exercises[i] = workout.Exercise{
			Name:        ex.name,
			MuscleGroup: ex.muscleGroup,
			Category:    ex.category,
			Sets:        ex.sets,
			RepsPerSet:  ex.repetitions,
			WeightKg:    ex.weight,
			Notes:       "",
		}
	}
	return exercises
}

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

func setDefaultWeight(weight float64) float64 {
	if weight < 0 {
		return 0
	}
	return weight
}

func setDefaultSets(sets int) int {
	if sets <= 0 {
		return 1
	}
	return sets
}

func setDefaultRepetitions(reps int) int {
	if reps <= 0 {
		return 1
	}
	return reps
}

func newCreateWorkoutDBReq(userID string, req database.CreateWorkoutReq) *database.CreateWorkoutReq {
	dbExercises := make([]createExerciseDBReq, len(req.GetExercises()))
	for i, ex := range req.GetExercises() {
		dbExercises[i] = createExerciseDBReq{
			name:        ex.GetName(),
			description: ex.GetDescription(),
			muscleGroup: constant.MuscleGroup(ex.GetMuscleGroup()),
			category:    constant.ExerciseCategory(ex.GetCategory()),
			repetitions: setDefaultRepetitions(ex.GetRepetitions()),
			sets:        setDefaultSets(ex.GetSets()),
			weight:      setDefaultWeight(ex.GetWeight()),
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

type workoutResponse struct {
	workout workout.Workout
}

func (r *workoutResponse) GetWorkout() workout.Workout {
	return r.workout
}

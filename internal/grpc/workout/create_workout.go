package workout

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
	"workout-training-api/internal/constant"
	workoutv1 "workout-training-api/internal/grpc/generated/workout-training-api/workout/v1"
	"workout-training-api/internal/postgres/db_types/workout"
)

func (w *Workout) CreateWorkout(ctx context.Context, req *workoutv1.CreateWorkoutRequest) (*workoutv1.CreateWorkoutResponse, error) {
	log := w.log.With("method", "CreateWorkout")

	if err := validateCreateWorkoutRequest(req); err != nil {
		log.ErrorContext(ctx, "invalid request", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	workout, err := w.ctrl.CreateWorkout(ctx, newCtrlReqCreateWorkout(req))
	if err != nil {
		log.ErrorContext(ctx, "failed to create workout", slog.Any("error", err))
		return nil, status.Error(codes.Internal, "failed to create workout")
	}

	log.InfoContext(
		ctx,
		"workout created successfully",
		slog.String("workout_id", workout.GetWorkout().WorkoutID),
	)

	return &workoutv1.CreateWorkoutResponse{
		Workout: &workoutv1.Workout{
			WorkoutId:     workout.GetWorkout().WorkoutID,
			Name:          req.GetName(),
			Description:   req.GetDescription(),
			Status:        req.GetStatus(),
			Exercises:     req.GetExercises(),
			ScheduledDate: req.GetScheduledDate(),
		},
	}, nil
}

type ctrlReqCreateWorkout struct {
	*workoutv1.CreateWorkoutRequest
}

func newCtrlReqCreateWorkout(req *workoutv1.CreateWorkoutRequest) *ctrlReqCreateWorkout {
	return &ctrlReqCreateWorkout{req}
}

func (w *ctrlReqCreateWorkout) GetExercises() []workout.Exercise {
	exercises := make([]workout.Exercise, 0, len(w.CreateWorkoutRequest.GetExercises()))
	for _, ex := range w.CreateWorkoutRequest.GetExercises() {
		exercises = append(exercises, workout.Exercise{
			ExerciseID:  ex.GetExerciseId(),
			Name:        ex.GetName(),
			Sets:        int(ex.GetSets()),
			RepsPerSet:  int(ex.GetRepsPerSet()),
			WeightKg:    float64(ex.GetWeightKg()),
			Notes:       ex.GetNotes(),
			MuscleGroup: constant.MuscleGroup(string(ex.GetMuscleGroup())),
			Category:    constant.ExerciseCategory(string(ex.GetCategory())),
		})
	}
	return exercises
}

func (w *ctrlReqCreateWorkout) GetStatus() constant.WorkoutStatus {
	return constant.WorkoutStatus(w.CreateWorkoutRequest.GetStatus())
}

func (w *ctrlReqCreateWorkout) GetScheduledDate() []time.Time {
	dates := make([]time.Time, 0, len(w.CreateWorkoutRequest.GetScheduledDate()))
	for _, d := range w.CreateWorkoutRequest.GetScheduledDate() {
		dates = append(dates, d.AsTime())
	}
	return dates
}

func (w *ctrlReqCreateWorkout) GetDescription() string {
	return w.CreateWorkoutRequest.GetDescription()
}

func validateCreateWorkoutRequest(req *workoutv1.CreateWorkoutRequest) error {
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "name is required")
	}
	if len(req.GetExercises()) == 0 {
		return status.Error(codes.InvalidArgument, "at least one exercise is required")
	}
	return nil
}

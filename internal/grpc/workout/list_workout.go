package workout

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	workoutv1 "workout-training-api/internal/grpc/generated/workout-training-api/workout/v1"
	"workout-training-api/internal/types/controller"
)

func (w *Workout) ListWorkout(ctx context.Context, req *workoutv1.ListWorkoutsRequest) (*workoutv1.ListWorkoutsResponse, error) {
	log := w.log.With("method", "ListWorkout")

	if err := validateListWorkoutRequest(req); err != nil {
		log.ErrorContext(ctx, "invalid request", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctrlReq := newCtrlReqListWorkout(req)

	ctrlResp, err := w.ctrl.ListWorkouts(ctx, ctrlReq)
	if err != nil {
		log.ErrorContext(ctx, "failed to list workouts", slog.Any("error", err))
		return nil, status.Error(codes.Internal, "failed to list workouts")
	}

	workouts := make([]*workoutv1.WorkoutView, 0, len(ctrlResp.GetWorkouts()))
	for _, workout := range ctrlResp.GetWorkouts() {
		workouts = append(workouts, convertToWorkoutView(workout))
	}

	log.InfoContext(ctx, "workouts listed successfully",
		slog.Int("count", len(workouts)))

	return &workoutv1.ListWorkoutsResponse{
		Workouts: workouts,
	}, nil
}

func convertToWorkoutView(workout controller.Workouts) *workoutv1.WorkoutView {
	exercises := make([]*workoutv1.Exercise, 0, len(workout.GetExercises()))
	for _, ex := range workout.GetExercises() {
		exercises = append(exercises, &workoutv1.Exercise{
			ExerciseId:  ex.GetID(),
			Name:        ex.GetName(),
			Description: ex.GetDescription(),
			MuscleGroup: workoutv1.MuscleGroup(workoutv1.MuscleGroup_value[ex.GetMuscleGroup()]),
			Category:    workoutv1.Category(workoutv1.Category_value[ex.GetCategory()]),
			RepsPerSet:  uint32(ex.GetRepetitions()),
			Sets:        uint32(ex.GetSets()),
			WeightKg:    float32(ex.GetWeight()),
		})
	}

	return &workoutv1.WorkoutView{
		WorkoutId:   workout.GetID(),
		Exercises:   exercises,
		WorkoutDate: timestamppb.New(workout.GetDate()),
		Comments:    workout.GetComments(),
	}
}

func validateListWorkoutRequest(req *workoutv1.ListWorkoutsRequest) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	if req.UserId == "" {
		return status.Error(codes.InvalidArgument, "user_id cannot be empty")
	}

	// Validate time range if provided
	if req.TimeRangeStart != nil && req.TimeRangeEnd != nil {
		if req.TimeRangeStart.AsTime().After(req.TimeRangeEnd.AsTime()) {
			return status.Error(codes.InvalidArgument, "time_range_start must be before time_range_end")
		}
	}

	return nil
}

type ctrlReqListWorkout struct {
	userID string
}

func (c *ctrlReqListWorkout) GetUserID() string {
	return c.userID
}

func newCtrlReqListWorkout(req *workoutv1.ListWorkoutsRequest) *ctrlReqListWorkout {
	return &ctrlReqListWorkout{
		userID: req.UserId,
	}
}

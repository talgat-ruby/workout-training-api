package workout

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	workoutv1 "workout-training-api/internal/grpc/generated/workout-training-api/workout/v1"
)

func (w *Workout) DeleteWorkout(ctx context.Context, req *workoutv1.DeleteWorkoutRequest) (*workoutv1.DeleteWorkoutResponse, error) {
	log := w.log.With("method", "DeleteWorkout")

	if err := validateDeleteWorkoutRequest(req); err != nil {
		log.ErrorContext(ctx, "invalid request", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := w.ctrl.DeleteWorkout(ctx, newCtrlReqDeleteWorkout(req)); err != nil {
		log.ErrorContext(ctx, "failed to delete workout", slog.Any("error", err))
		return nil, status.Error(codes.Internal, "failed to delete workout")
	}

	log.InfoContext(
		ctx,
		"workout deleted successfully",
		slog.String("workout_id", req.GetWorkoutId()),
	)

	return &workoutv1.DeleteWorkoutResponse{}, nil
}

func validateDeleteWorkoutRequest(req *workoutv1.DeleteWorkoutRequest) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	if req.WorkoutId == "" {
		return status.Error(codes.InvalidArgument, "workout_id cannot be empty")
	}

	return nil
}

type ctrlReqDeleteWorkout struct {
	id string
}

func (c *ctrlReqDeleteWorkout) GetID() string {
	return c.id
}

func newCtrlReqDeleteWorkout(req *workoutv1.DeleteWorkoutRequest) *ctrlReqDeleteWorkout {
	return &ctrlReqDeleteWorkout{
		id: req.WorkoutId,
	}
}

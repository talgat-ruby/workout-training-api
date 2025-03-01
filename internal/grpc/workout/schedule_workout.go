package workout

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
	workoutv1 "workout-training-api/internal/grpc/generated/workout-training-api/workout/v1"
)

func (w *Workout) ScheduleWorkout(ctx context.Context, req *workoutv1.ScheduleWorkoutRequest) (*workoutv1.ScheduleWorkoutResponse, error) {
	log := w.log.With("method", "ScheduleWorkout")

	if err := validateScheduleWorkoutRequest(req); err != nil {
		log.ErrorContext(ctx, "invalid request", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := w.ctrl.ScheduleWorkout(ctx, newCtrlReqScheduleWorkout(req)); err != nil {
		log.ErrorContext(ctx, "failed to schedule workout", slog.Any("error", err))
		return nil, status.Error(codes.Internal, "failed to schedule workout")
	}

	log.InfoContext(
		ctx,
		"workout scheduled successfully",
		slog.String("workout_id", req.GetWorkoutId()),
		slog.Time("scheduled_date", req.GetScheduledDate().AsTime()),
	)

	return &workoutv1.ScheduleWorkoutResponse{}, nil
}

func validateScheduleWorkoutRequest(req *workoutv1.ScheduleWorkoutRequest) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	if req.WorkoutId == "" {
		return status.Error(codes.InvalidArgument, "workout_id cannot be empty")
	}

	if req.ScheduledDate == nil {
		return status.Error(codes.InvalidArgument, "scheduled_date cannot be nil")
	}

	scheduledDate := req.ScheduledDate.AsTime()
	if scheduledDate.Before(time.Now()) {
		return status.Error(codes.InvalidArgument, "scheduled_date cannot be in the past")
	}

	return nil
}

type ctrlReqScheduleWorkout struct {
	workoutID string
	date      time.Time
}

func (c *ctrlReqScheduleWorkout) GetWorkoutID() string {
	return c.workoutID
}

func (c *ctrlReqScheduleWorkout) GetDate() time.Time {
	return c.date
}

func newCtrlReqScheduleWorkout(req *workoutv1.ScheduleWorkoutRequest) *ctrlReqScheduleWorkout {
	return &ctrlReqScheduleWorkout{
		workoutID: req.WorkoutId,
		date:      req.ScheduledDate.AsTime(),
	}
}

package workout

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	workoutv1 "workout-training-api/internal/grpc/generated/workout-training-api/workout/v1"
	"workout-training-api/internal/types/controller"
)

func (w *Workout) UpdateWorkout(ctx context.Context, req *workoutv1.UpdateWorkoutRequest) (*workoutv1.UpdateWorkoutResponse, error) {
	log := w.log.With("method", "UpdateWorkout")

	if err := validateUpdateWorkoutRequest(req); err != nil {
		log.ErrorContext(ctx, "invalid request", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := w.ctrl.UpdateWorkout(ctx, newCtrlReqUpdateWorkout(req)); err != nil {
		log.ErrorContext(ctx, "failed to update workout", slog.Any("error", err))
		return nil, status.Error(codes.Internal, "failed to update workout")
	}

	log.InfoContext(
		ctx,
		"workout updated successfully",
		slog.String("workout_id", req.GetWorkoutId()),
	)

	return &workoutv1.UpdateWorkoutResponse{}, nil
}

func validateUpdateWorkoutRequest(req *workoutv1.UpdateWorkoutRequest) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	if req.WorkoutId == "" {
		return status.Error(codes.InvalidArgument, "workout_id cannot be empty")
	}

	if len(req.Exercises) == 0 {
		return status.Error(codes.InvalidArgument, "exercises cannot be empty")
	}

	for _, exercise := range req.Exercises {
		if exercise == nil {
			return status.Error(codes.InvalidArgument, "exercise cannot be nil")
		}
		if exercise.Name == "" {
			return status.Error(codes.InvalidArgument, "exercise name cannot be empty")
		}
		if exercise.Sets <= 0 {
			return status.Error(codes.InvalidArgument, "exercise sets must be greater than 0")
		}
		if exercise.RepsPerSet <= 0 {
			return status.Error(codes.InvalidArgument, "exercise repetitions must be greater than 0")
		}
		if exercise.WeightKg < 0 {
			return status.Error(codes.InvalidArgument, "exercise weight cannot be negative")
		}
		if exercise.MuscleGroup == workoutv1.MuscleGroup_MUSCLE_GROUP_UNSPECIFIED {
			return status.Error(codes.InvalidArgument, "exercise muscle group must be specified")
		}
		if exercise.Category == workoutv1.Category_CATEGORY_UNSPECIFIED {
			return status.Error(codes.InvalidArgument, "exercise category must be specified")
		}
	}

	return nil
}

type ctrlReqUpdateWorkout struct {
	id        string
	exercises []controller.Exercise
	comments  []string
}

func (c *ctrlReqUpdateWorkout) GetID() string {
	return c.id
}

func (c *ctrlReqUpdateWorkout) GetExercises() []controller.Exercise {
	return c.exercises
}

func (c *ctrlReqUpdateWorkout) GetComments() []string {
	return c.comments
}

func newCtrlReqUpdateWorkout(req *workoutv1.UpdateWorkoutRequest) *ctrlReqUpdateWorkout {
	exercises := make([]controller.Exercise, 0, len(req.Exercises))
	for _, ex := range req.Exercises {
		exercises = append(exercises, &Exercise{
			ID:          ex.ExerciseId,
			Name:        ex.Name,
			Description: ex.Description,
			MuscleGroup: string(ex.MuscleGroup),
			Category:    string(ex.Category),
			Repetitions: int(ex.RepsPerSet),
			Sets:        int(ex.Sets),
			Weight:      float64(ex.WeightKg),
		})
	}

	return &ctrlReqUpdateWorkout{
		id:        req.WorkoutId,
		exercises: exercises,
		comments:  req.Comments,
	}
}

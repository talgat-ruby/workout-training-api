package graph

import (
	"context"
	"log/slog"
	"time"
	"workout-training-api/internal/graphql/graph/model"
	"workout-training-api/pkg/httputils/request"
)

func (r *mutationResolver) CreateWorkout(ctx context.Context, name string, description string, exercises []model.Exercise, scheduledTime time.Time) (bool, error) {
	log := r.logger.With("method", "CreateWorkout")

	reqBody := request.NewCtrlCreateWorkoutRequest(name, description, exercises, scheduledTime)
	_, err := r.ctrl.CreateWorkout(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		return false, err
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return true, nil
}

func (r *mutationResolver) DeleteWorkout(ctx context.Context, workoutID string) (bool, error) {
	log := r.logger.With("method", "CreateWorkout")

	reqBody := request.NewCtrlDeleteWorkoutRequest(workoutID)
	_, err := r.ctrl.DeleteWorkout(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		return false, err
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return true, nil
}

func (r *mutationResolver) ListWorkouts(ctx context.Context) (model.Workout, error) {
	log := r.logger.With("method", "CreateWorkout")

	list, err := r.ctrl.ListWorkouts(ctx)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		return model.Workout{}, err
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return list, nil
}

// TODO controller method?
//func (r *mutationResolver) AddExercise(ctx context.Context, workoutId string, exerciseId string) (bool, error) {
//	log := r.logger.With("method", "AddExercise")
//
//	reqBody := request.NewCtrlAddExerciseRequest(workoutId, exerciseId)
//	_, err := r.ctrl.AddExercise(ctx, reqBody)
//	if err != nil {
//		log.ErrorContext(ctx, "fail", slog.Any("error", err))
//		return false, err
//	}
//
//	log.InfoContext(
//		ctx,
//		"success",
//	)
//	return true, nil
//}

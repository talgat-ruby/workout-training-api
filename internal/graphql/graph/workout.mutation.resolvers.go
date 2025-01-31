package graph

import (
	"context"
	"log/slog"
)

// CreateWorkout TODO return generated workout ID
func (r *mutationResolver) CreateWorkout(ctx context.Context, name string) (bool, error) {
	log := r.logger.With("method", "CreateWorkout")

	reqBody := request.NewCtrlCreateWorkoutRequest(name)
	_, err := r.ctrl.NewWorkout(ctx, reqBody)
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

// AddExercise adds exercises into workout.Exercises
func (r *mutationResolver) AddExercise(ctx context.Context, workoutId string, exerciseId string) (bool, error) {
	log := r.logger.With("method", "AddExercise")

	// TODO id is string or number?
	reqBody := request.NewCtrlAddExerciseRequest(workoutId, exerciseId)
	_, err := r.ctrl.AddExercise(ctx, reqBody)
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

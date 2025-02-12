package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"log/slog"
	"time"
	"workout-training-api/internal/graphql/graph/model"
	"workout-training-api/pkg/httputils/request"
)

// CreateWorkout is the resolver for the createWorkout field.
func (r *mutationResolver) CreateWorkout(ctx context.Context, name string, description string, exercises []*model.ExerciseInput, scheduledTime time.Time) (bool, error) {
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

// DeleteWorkout is the resolver for the deleteWorkout field.
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

// UpdateWorkout is the resolver for the updateWorkout field.
func (r *mutationResolver) UpdateWorkout(ctx context.Context, workout *model.WorkoutInput) (model.WorkoutInput, error) {
	log := r.logger.With("method", "UpdateWorkout")

	reqBody := request.NewCtrlUpdateWorkoutRequest(workout)
	_, err := r.ctrl.UpdateWorkout(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		return *workout, err
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return *workout, nil
}

// ListWorkouts is the resolver for the listWorkouts field.
func (r *queryResolver) ListWorkouts(ctx context.Context) ([]*model.Workout, error) {
	log := r.logger.With("method", "ListWorkouts")

	list, err := r.ctrl.ListWorkouts(ctx)

	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		return nil, err
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return list, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *exerciseResolver) Repetitions(ctx context.Context, obj *model.Exercise) (int32, error) {
	panic(fmt.Errorf("not implemented: Repetitions - repetitions"))
}
func (r *exerciseResolver) Sets(ctx context.Context, obj *model.Exercise) (int32, error) {
	panic(fmt.Errorf("not implemented: Sets - sets"))
}
func (r *exerciseResolver) Weight(ctx context.Context, obj *model.Exercise) (float64, error) {
	panic(fmt.Errorf("not implemented: Weight - weight"))
}
func (r *workoutResolver) ScheduledTimes(ctx context.Context, obj *model.Workout) (int32, error) {
	panic(fmt.Errorf("not implemented: ScheduledTimes - scheduledTimes"))
}
func (r *Resolver) Exercise() ExerciseResolver { return &exerciseResolver{r} }
func (r *Resolver) Workout() WorkoutResolver { return &workoutResolver{r} }
type exerciseResolver struct{ *Resolver }
type workoutResolver struct{ *Resolver }
*/

package workout

import (
	"context"
	"fmt"
	"time"
	"workout-training-api/internal/types/controller"
	"workout-training-api/internal/types/database"
)

type WorkoutController struct {
	db database.Workout
}

func (w *WorkoutController) CreateWorkout(ctx context.Context, req controller.CreateWorkoutReq) (controller.CreateWorkoutResp, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, fmt.Errorf("user not authenticated")
	}

	exerciseIDs := make([]string, 0, len(req.GetExercises()))
	for _, e := range req.GetExercises() {
		exerciseIDs = append(exerciseIDs, e.GetID())
	}

	times := make([]time.Time, 0, len(req.GetDate()))
	for _, t := range req.GetDate() {
		times = append(times, t)
	}

	workout := &Workout{
		UserID:         userID,
		Exercises:      exerciseIDs,
		ScheduledTimes: times,
	}

	_, err := w.db.CreateWorkout(ctx, workout)
	if err != nil {
		return nil, fmt.Errorf("failed to save workout: %w", err)
	}

	return &createWorkoutResponse{}, nil
}

type createWorkoutResponse struct {
}

type Workout struct {
	ID             string
	UserID         string
	Exercises      []string
	Comments       []string
	ScheduledTimes []time.Time
}

func (w *Workout) GetUserID() string {
	return w.UserID
}

func (w *Workout) GetExerciseIDs() []string {
	return w.Exercises
}

func (w *Workout) GetScheduledTimes() []time.Time {
	return w.ScheduledTimes
}

func (w *Workout) GetID() string {
	return w.ID
}

func (w *Workout) GetComments() []string {
	return w.Comments
}

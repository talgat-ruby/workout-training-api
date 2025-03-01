package workout

import (
	"context"
	"fmt"
	"time"
	"workout-training-api/internal/constant"
	"workout-training-api/internal/types/controller"
)

func (w *WorkoutController) UpdateWorkout(ctx context.Context, req controller.UpdateWorkoutReq) (controller.UpdateWorkoutResp, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, fmt.Errorf("user not authenticated")
	}

	exerciseIDs := make([]string, 0, len(req.GetExercises()))
	for _, e := range req.GetExercises() {
		exerciseIDs = append(exerciseIDs, e.GetID())
	}

	workout := &WorkoutUP{
		ID:        req.GetID(),
		UserID:    userID,
		Exercises: exerciseIDs,
		Comments:  req.GetComments(),
	}

	_, err := w.db.UpdateWorkout(ctx, workout)
	if err != nil {
		return nil, fmt.Errorf("failed to update workout: %w", err)
	}

	return &updateWorkoutResponse{}, nil
}

type WorkoutUP struct {
	ID            string
	UserID        string
	Name          string
	Description   string
	Status        constant.WorkoutStatus
	Exercises     []string
	Comments      []string
	ScheduledDate time.Time
}

type updateWorkoutResponse struct{}

func (w *WorkoutUP) GetComments() []string {
	return w.Comments
}

func (w *WorkoutUP) GetExerciseIDs() []string {
	return w.Exercises
}

func (w *WorkoutUP) GetID() string {
	return w.ID
}

func (w *WorkoutUP) GetUserID() string {
	return w.UserID
}

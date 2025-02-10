package request

import (
	"time"
	"workout-training-api/internal/graphql/graph/model"
)

type CtrlCreateWorkoutRequest struct {
	name          string
	description   string
	exercises     []*model.ExerciseInput
	scheduledTime time.Time
}

func NewCtrlCreateWorkoutRequest(name string, description string, exercises []*model.ExerciseInput, scheduledTime time.Time) *CtrlCreateWorkoutRequest {
	return &CtrlCreateWorkoutRequest{
		name: name, description: description, exercises: exercises, scheduledTime: scheduledTime,
	}
}

func NewCtrlDeleteWorkoutRequest(workoutID string) string {
	return workoutID
}

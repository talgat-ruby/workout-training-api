package workout

import (
	"context"

	"workout-training-api/internal/types/database"
)

func (m *Workout) CreateExercise(ctx context.Context, req database.CreateExerciseReq) (database.CreateExerciseResp, error) {
	panic("implement me")
}


func (m *Workout) UpdateExercise(ctx context.Context, req  database.UpdateExerciseReq) (database.UpdateExerciseResp, error) {

}

func (m *Workout) DeleteExercise(ctx context.Context, req database.DeleteExerciseReq) (database.DeleteExerciseResp, error)

func (m *Workout) ListExpense(ctx context.Context, req database.) (database.ListWorkoutResp, error)
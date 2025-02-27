package exercise

import (
	"context"
	"fmt"
	"time"
	"workout-training-api/internal/types/database"
)

func (m *ExerciseModel) DeleteExercise(ctx context.Context, req database.DeleteExerciseReq) (database.DeleteExerciseResp, error) {
	log := m.logger.With("exercise", "DeleteExercise")

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	query := `
		DELETE FROM exercises
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := m.db.ExecContext(ctx, query, req.GetID())
	if err != nil {
		log.ErrorContext(ctx, "failed to delete exercise", "id", req.GetID())
		return nil, fmt.Errorf("failed to delete exercise: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "failed to call rowsAffected", "id", req.GetID())
		return nil, fmt.Errorf("failed to call rowsAffected: %w", err)
	}

	if rowsAffected == 0 {
		log.ErrorContext(ctx, "failed to find an exercise with id", "id", req.GetID())
		return nil, fmt.Errorf("failed to find an exercise")
	}

	return true, nil
}

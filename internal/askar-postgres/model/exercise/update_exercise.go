package exercise

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"time"
	"workout-training-api/internal/types/database"
)

func (m *ExerciseModel) UpdateExercise(ctx context.Context, req database.UpdateExerciseReq) (database.UpdateExerciseResp, error) {
	log := m.logger.With("exercise", "UpdateExercise")

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	query := `
		UPDATE exercises
		SET name = $1, description = $2, categories = $3, muscle_groups = $4, updated_at = now()
		WHERE id = $5`

	args := []interface{}{
		req.GetName(),
		req.GetDescription(),
		pq.Array(req.GetCategories()),
		pq.Array(req.GetMuscleGroups()),
		req.GetID(),
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.ErrorContext(ctx, "failed to update exercise", "error", err)
		return nil, fmt.Errorf("failed to update exercise: %w", err)
	}

	return true, nil
}

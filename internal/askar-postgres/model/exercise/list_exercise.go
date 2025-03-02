package exercise

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"time"
	"workout-training-api/internal/types/database"
)

type ExerciseList struct {
	exercises []database.ExerciseResp
}

func (e ExerciseList) GetList() []database.ExerciseResp {
	return e.exercises
}

func (m *ExerciseModel) ListExercise(ctx context.Context, _ database.ListExerciseReq) (database.ListExerciseResp, error) {
	log := m.logger.With("exercise", "ListExercise")

	query := `
		SELECT id, name, description, repetitions, sets, weight, muscle_groups, categories, created_at, updated_at 
		FROM exercises`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		log.ErrorContext(ctx, "failed to query exercises")
		return nil, fmt.Errorf("failed to query exercises: %w", err)
	}

	exercises := []database.ExerciseResp{}

	for rows.Next() {
		var exercise Exercise

		err := rows.Scan(
			&exercise.ID,
			&exercise.Name,
			&exercise.Description,
			&exercise.Repetitions,
			&exercise.Sets,
			&exercise.Weight,
			pq.Array(&exercise.MuscleGroups),
			pq.Array(&exercise.Categories),
			&exercise.CreatedAt,
			&exercise.UpdatedAt,
		)
		if err != nil {
			log.ErrorContext(ctx, "failed to scan exercise")
			return nil, fmt.Errorf("failed to scan exercise: %w", err)
		}

		exercises = append(exercises, &exercise)
	}

	return ExerciseList{exercises: exercises}, nil
}

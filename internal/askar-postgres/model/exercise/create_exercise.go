package exercise

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"time"
	"workout-training-api/internal/types/database"
)

func (m *ExerciseModel) CreateExercise(ctx context.Context, req database.CreateExerciseReq) (database.CreateExerciseResp, error) {

	log := m.logger.With("exercise", "CreateExercise")

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	mdl := &Exercise{
		Name:         req.GetName(),
		Description:  req.GetDescription(),
		Repetitions:  req.GetRepetitions(),
		Sets:         req.GetSets(),
		Weight:       req.GetWeight(),
		MuscleGroups: req.GetMuscleGroups(),
		Categories:   req.GetCategories(),
	}

	query := `
		INSERT INTO exercises (name, description, repetitions, sets, weight, muscle_groups, categories)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	args := []interface{}{
		mdl.Name,
		mdl.Description,
		mdl.Repetitions,
		mdl.Sets,
		mdl.Weight,
		pq.Array(mdl.MuscleGroups),
		pq.Array(mdl.Categories),
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.ErrorContext(ctx, "failed to execute query", "error", err)
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return true, nil
}

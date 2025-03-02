package seeds

import (
	"context"
	"fmt"
	"log/slog"
	askar_postgres "workout-training-api/internal/askar-postgres"
	"workout-training-api/internal/askar-postgres/model/exercise"
)

type Seed interface {
	Populate() error
	TestMethods() error
}

type seeder struct {
	postgres *askar_postgres.Postgres
}

func New(postgres *askar_postgres.Postgres) Seed {
	return &seeder{postgres: postgres}
}

func (s *seeder) Populate() error {

	exercises := []*exercise.Exercise{
		{
			Name:         "Push-ups",
			Description:  "A bodyweight exercise that targets the chest, shoulders, and triceps.",
			Repetitions:  15,
			Sets:         3,
			Weight:       0,
			MuscleGroups: []string{"Chest", "Triceps", "Shoulders"},
			Categories:   []string{"Bodyweight", "Strength"},
		},
		{
			Name:         "Squats",
			Description:  "A compound exercise that targets the legs and glutes.",
			Repetitions:  12,
			Sets:         3,
			Weight:       50,
			MuscleGroups: []string{"Quadriceps", "Hamstrings", "Glutes"},
			Categories:   []string{"Legs", "Strength"},
		},
		{
			Name:         "Deadlifts",
			Description:  "A strength exercise that targets the back, legs, and core.",
			Repetitions:  8,
			Sets:         4,
			Weight:       100,
			MuscleGroups: []string{"Back", "Hamstrings", "Core"},
			Categories:   []string{"Powerlifting", "Strength"},
		},
	}

	for _, exercise := range exercises {
		_, err := s.postgres.ExerciseModel.CreateExercise(context.Background(), exercise)
		if err != nil {
			slog.Error("Error on create exercise", err.Error())
		}
	}

	return nil
}

func (s *seeder) TestMethods() error {
	s.listExercise()
	return nil
}

func (s *seeder) deleteExercise() {
	req := &exercise.Exercise{
		ID: 1,
	}
	_, err := s.postgres.DeleteExercise(context.Background(), req)
	if err != nil {
		slog.Error("Error on deleting exercise", err.Error())
	}
}

func (s *seeder) listExercise() {
	res, err := s.postgres.ExerciseModel.ListExercise(context.Background(), nil)
	if err != nil {
		slog.Error("Error on list exercise", err.Error())
	}

	for _, exercise := range res.GetList() {
		fmt.Printf("exercise id %s and name %s\n", exercise.GetID(), exercise.GetName())
	}
}

func (s *seeder) updateExercise() {
	exercise := &exercise.Exercise{
		ID:           2,
		Name:         "Updated Squats",
		Description:  "Updated A compound exercise that targets the legs and glutes.",
		MuscleGroups: []string{"Updated Quadriceps", "Hamstrings", "Glutes"},
		Categories:   []string{"Updated Legs", "Strength"},
	}

	_, err := s.postgres.UpdateExercise(context.Background(), exercise)
	if err != nil {
		slog.Error("Error on update exercise", err.Error())
	}
	slog.Info("Exercise updated")
}

package exercise

import (
	"strconv"
	"time"
)

type Exercise struct {
	ID           int64
	Name         string
	Description  string
	Repetitions  int32
	Sets         int32
	Weight       int32
	MuscleGroups []string
	Categories   []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (e *Exercise) GetWorkoutID() string {
	return "id"
}

func (e *Exercise) GetID() string {
	return strconv.FormatInt(e.ID, 10)
}

func (e *Exercise) GetName() string {
	return e.Name
}

func (e *Exercise) GetDescription() string {
	return e.Description
}

func (e *Exercise) GetRepetitions() int32 { return e.Repetitions }

func (e *Exercise) GetSets() int32 { return e.Sets }

func (e *Exercise) GetWeight() int32 { return e.Weight }

func (e *Exercise) GetMuscleGroups() []string {
	return e.MuscleGroups
}

func (e *Exercise) GetCategories() []string {
	return e.Categories
}

func (e *Exercise) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *Exercise) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

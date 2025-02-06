package database

type CreateExerciseReq interface {
	GetWorkoutID() string
	GetName() string
	GetDescription() string
	GetCategory() string
	GetMuscleGroup() string
	GetSets() int
	GetRepsPerSet() int
	GetWeightKg() float64
	GetNotes() string
}

type CreateExerciseResp interface {
}

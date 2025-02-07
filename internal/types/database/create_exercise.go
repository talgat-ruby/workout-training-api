package database

type CreateExerciseReq interface {
	GetWorkoutID() string
	GetName() string
	GetCategory() string
	GetMuscleGroup() string
	GetSets() int
	GetRepsPerSet() int
	GetWeightKg() float64
	GetNotes() string
}

type CreateExerciseResp interface {
}

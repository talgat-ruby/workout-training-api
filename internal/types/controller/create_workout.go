package controller

type CreateWorkoutReq interface {
    GetExercises() []Exercise
    GetDate() string
}

type CreateWorkoutResp interface {
    GetID() string
}

type Exercise interface {
    GetName() string
    GetDescription() string
    GetCategory() string
    GetRepetitions() int
    GetSets() int
    GetWeight() float64
}


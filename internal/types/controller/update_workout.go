package controller

type UpdateWorkoutReq interface {
    GetID() string
    GetExercises() []Exercise
    GetComments() string
}

type UpdateWorkoutResp interface{}


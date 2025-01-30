package controller

import (
    "time"
)

type ListWorkoutsReq interface {
    GetUserID() string
    GetStatus() string
}

type ListWorkoutsResp interface {
    GetWorkouts() []Workout
}

type Workout interface {
    GetID() string
    GetExercises() []Exercise
    GetDate() string
    GetComments() string
}

type Exercise interface {
    GetName() string
    GetDescription() string
    GetCategory() string
    GetRepetitions() int
    GetSets() int
    GetWeight() float64
}


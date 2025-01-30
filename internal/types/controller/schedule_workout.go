package controller

type ScheduleWorkoutReq interface {
    GetWorkoutID() string
    GetDate() string
}

type ScheduleWorkoutResp interface{}


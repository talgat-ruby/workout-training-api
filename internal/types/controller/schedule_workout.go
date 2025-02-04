package controller

import "time"

type ScheduleWorkoutReq interface {
	GetWorkoutID() string
	GetDate() time.Time
}

type ScheduleWorkoutResp interface{}

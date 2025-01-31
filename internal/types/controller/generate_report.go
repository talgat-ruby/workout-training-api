package controller

import "time"

type GenerateReportReq interface {
	GetUserID() string
	GetStartDate() time.Time
	GetEndDate() time.Time
}

type GenerateReportResp interface {
	GetReport() string
}

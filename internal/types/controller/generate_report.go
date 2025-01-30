package controller

type GenerateReportReq interface {
    GetUserID() string
    GetStartDate() string
    GetEndDate() string
}

type GenerateReportResp interface {
    GetReport() string
}


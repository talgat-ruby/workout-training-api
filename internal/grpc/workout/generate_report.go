package workout

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
	workoutv1 "workout-training-api/internal/grpc/generated/workout-training-api/workout/v1"
)

func (w *Workout) GenerateReport(ctx context.Context, req *workoutv1.GenerateReportRequest) (*workoutv1.GenerateReportResponse, error) {
	log := w.log.With("method", "GenerateReport")

	if err := validateGenerateRequest(req); err != nil {
		log.ErrorContext(ctx, "invalid request", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	report, err := w.ctrl.GenerateReport(ctx, newCtrlReqGenerateReport(req))
	if err != nil {
		log.ErrorContext(ctx, "failed to generate report", slog.Any("error", err))
		return nil, status.Error(codes.Internal, "failed to generate workout report")
	}

	log.InfoContext(
		ctx,
		"report generated successfully",
		slog.String("user_id", req.GetUserId()),
		slog.Time("start_date", req.GetStartDate().AsTime()),
		slog.Time("end_date", req.GetEndDate().AsTime()),
	)

	return &workoutv1.GenerateReportResponse{
		Report: report.GetReport(),
	}, nil
}

func validateGenerateRequest(req *workoutv1.GenerateReportRequest) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	if req.UserId == "" {
		return status.Error(codes.InvalidArgument, "user_id cannot be empty")
	}

	startDate := req.GetStartDate().AsTime()
	endDate := req.GetEndDate().AsTime()

	if startDate.IsZero() {
		return status.Error(codes.InvalidArgument, "start_date cannot be empty")
	}

	if endDate.IsZero() {
		return status.Error(codes.InvalidArgument, "end_date cannot be empty")
	}

	if endDate.Before(startDate) {
		return status.Error(codes.InvalidArgument, "end_date cannot be before start_date")
	}

	return nil
}

type ctrlGenerateReportReq struct {
	userID    string
	startDate time.Time
	endDate   time.Time
}

func (r *ctrlGenerateReportReq) GetUserID() string {
	return r.userID
}

func (r *ctrlGenerateReportReq) GetStartDate() time.Time {
	return r.startDate
}

func (r *ctrlGenerateReportReq) GetEndDate() time.Time {
	return r.endDate
}

func newCtrlReqGenerateReport(req *workoutv1.GenerateReportRequest) *ctrlGenerateReportReq {
	return &ctrlGenerateReportReq{
		userID:    req.GetUserId(),
		startDate: req.GetStartDate().AsTime(),
		endDate:   req.GetEndDate().AsTime(),
	}
}

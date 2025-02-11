package workout

//import (
//	"context"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//	"log/slog"
//
//	workoutv1 "workout-training-api/internal/grpc/generated/workout-training-api/workout/v1"
//)
//
//func (w *Workout) ListWorkout(ctx context.Context, req *workoutv1.ListWorkoutsRequest) (*workoutv1.ListWorkoutsResponse, error) {
//	log := w.log.With("method", "ListWorkout")
//
//	if err := validateListWorkoutRequest(req); err != nil {
//		log.ErrorContext(ctx, "invalid request", slog.Any("error", err))
//		return nil, status.Error(codes.InvalidArgument, err.Error())
//	}
//
//	ctrlReq := newCtrlReqListWorkout(req)
//
//	ctrlResp, err := w.ctrl.ListWorkouts(ctx, ctrlReq)
//	if err != nil {
//		log.ErrorContext(ctx, "failed to list workouts", slog.Any("error", err))
//		return nil, status.Error(codes.Internal, "failed to list workouts")
//	}
//
//	log.InfoContext(ctx, "workouts listed successfully")
//
//	return &workoutv1.ListWorkoutsResponse{
//		Workouts: ctrlResp.GetWorkouts(),
//	}, nil
//}
//
//func validateListWorkoutRequest(req *workoutv1.ListWorkoutsRequest) error {
//	if req == nil {
//		return status.Error(codes.InvalidArgument, "request cannot be nil")
//	}
//
//	return nil
//}
//
//type ctrlReqListWorkout struct {
//	userID string
//}
//
//func (c ctrlReqListWorkout) GetUserID() string {
//	return c.userID
//}
//
//func newCtrlReqListWorkout(req *workoutv1.ListWorkoutsRequest) *ctrlReqListWorkout {
//	return &ctrlReqListWorkout{
//		userID: req.UserId,
//	}
//}

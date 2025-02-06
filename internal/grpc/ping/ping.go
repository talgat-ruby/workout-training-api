package ping

import (
	"context"
	"fmt"
	pingSrvc "workout-training-api/internal/grpc/generated/workout-training-api/ping/v1"
)

func (s *Ping) Ping(_ context.Context, req *pingSrvc.PingRequest) (*pingSrvc.PingResponse, error) {
	return &pingSrvc.PingResponse{
		Message: fmt.Sprint("pong: ", req.GetMessage()),
	}, nil

}

package ping

import "fmt"

func (s *Sanitary) Ping(_ context.Context, req *sanitaryv1.PingRequest) (*sanitaryv1.PingResponse, error) {
	return &sanitaryv1.PingResponse{
		Message: fmt.Sprint("pong: ", req.GetMessage()),
	}, nil
}

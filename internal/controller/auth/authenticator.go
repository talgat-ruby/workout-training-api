package auth

import (
	"context"
)

func (c AuthController) Authenticator(ctx context.Context, str string) (context.Context, error) {
	panic("implement me")
}

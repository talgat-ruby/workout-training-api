package database

import "context"

type CreateUserReq interface {
	GetEmail() string
	GetPasswordHash() string
}

type CreateUserResp interface {
	GetID() string
}

type createUserResp struct {
	ID string
}

func (r createUserResp) GetID() string {
	return r.ID
}

func (d *Database) CreateUser(ctx context.Context, req CreateUserReq) (CreateUserResp, error) {
	var id string
	err := d.DB.QueryRowContext(ctx, `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2) RETURNING id`,
		req.GetEmail(), req.GetPasswordHash(),
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return createUserResp{ID: id}, nil
}

package database

<<<<<<< HEAD
import "context"
=======
import "time"
>>>>>>> 12012e7eedb63c6cad288001c5cd429cb4ec7dde

type FindUserReq interface {
	GetEmail() string
}

type FindUserResp interface {
	GetID() string
	GetEmail() string
	GetPasswordHash() string
	GetSalt() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type findUserResp struct {
	ID    string
	Email string
}

func (r findUserResp) GetID() string {
	return r.ID
}

func (r findUserResp) GetEmail() string {
	return r.Email
}

func (d *Database) FindUser(ctx context.Context, req FindUserReq) (FindUserResp, error) {
	var id, email string
	err := d.DB.QueryRowContext(ctx, `
		SELECT id, email FROM users WHERE email = $1 AND password_hash = $2`,
		req.GetEmail(), req.GetPasswordHash(),
	).Scan(&id, &email)

	if err != nil {
		return nil, err
	}

	return findUserResp{ID: id, Email: email}, nil
}

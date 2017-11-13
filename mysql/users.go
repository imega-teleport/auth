package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/imega-teleport/auth/api"
	"github.com/imega-teleport/auth/model"
	"golang.org/x/net/context"
)

func (r *repo) GetUser(ctx context.Context, login, pass string) (*auth.User, error) {
	var u auth.User
	row := r.db.QueryRowContext(ctx, "SELECT login, pass, created_at, active FROM users WHERE login = ? and pass = ?", login, pass)
	err := row.Scan(&u.Login, &u.Pass, &u.CreateAt, &u.Active)
	switch {
	case err == sql.ErrNoRows:
		return &auth.User{}, errors.New("user not exists")
	case err != nil:
		return &auth.User{}, fmt.Errorf("fail get user: %s", err.Error())
	}

	return &u, nil
}

// NewRepository returns new instance.
func NewRepository(opts ...Option) model.Repository {
	r := &repo{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

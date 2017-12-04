package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
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

func (r *repo) AuthUser(ctx context.Context, login, pass string) error {
	curTime := time.Now().Format("2006-01-02 15:04:05")
	var u auth.User
	row := r.db.QueryRowContext(ctx, "SELECT login, pass, created_at, active FROM users WHERE login = ? and pass = ? and active = 1 and expired_at >= ? limit 1", login, pass, curTime)
	err := row.Scan(&u.Login, &u.Pass, &u.CreateAt, &u.Active)
	switch {
	case err == sql.ErrNoRows:
		return errors.New("user not exists")
	case err != nil:
		return fmt.Errorf("failed auth user: %s", err.Error())
	}

	return nil
}

func (r *repo) CreateUser(ctx context.Context, user *auth.User) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed start transaction on create a data of user %s", err)
	}
	defer tx.ErrorHandle(&err, "create a data of user")
	return tx.createUser(ctx, user)
}

func (tx *Tx) createUser(ctx context.Context, user *auth.User) error {
	timeCreate, err := time.Parse("2006-01-02 15:04:05", user.GetCreateAt())
	if err != nil {
		return fmt.Errorf("failed convert create date from string %s", err)
	}

	b := squirrel.Insert("users").Columns("login", "pass", "created_at", "expired_at", "active")
	q, args, err := b.Values(user.GetLogin(), user.GetPass(), timeCreate, timeCreate, user.GetActive()).ToSql()
	if err != nil {
		return fmt.Errorf("failed builded query %s", err)
	}

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed exec a query on create a data of user %s", err)
	}
	return nil
}

// NewRepository returns new instance.
func NewRepository(opts ...Option) model.Repository {
	r := &repo{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

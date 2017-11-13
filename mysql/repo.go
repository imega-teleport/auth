package mysql

import (
	"database/sql"
	"fmt"

	"golang.org/x/net/context"
)

// Option is any options for package
type Option func(r *repo)

// WithDB used sql package
func WithDB(sqlDB *sql.DB) Option {
	return func(r *repo) {
		r.db = db{sqlDB}
	}
}

type repo struct {
	db db
}

type db struct {
	*sql.DB
}

func (d db) Begin(ctx context.Context) (*Tx, error) {
	tx, err := d.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Tx{tx}, nil
}

// Tx is a transaction
type Tx struct {
	*sql.Tx
}

// ErrorHandle closes yours opened any transaction
func (tx Tx) ErrorHandle(err *error, strLog string) {
	if *err != nil {
		if errRol := tx.Rollback(); errRol != nil {
			*err = fmt.Errorf("Failed rollback transaction after exec a query on %s %s", strLog, errRol)
		}
		return
	}

	*err = tx.Commit()
	if *err != nil {
		*err = fmt.Errorf("Failed commit transaction on %s %s", strLog, *err)
	}
}

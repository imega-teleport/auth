package model

import (
	"github.com/imega-teleport/auth/api"
	"golang.org/x/net/context"
)

// Repository represents example repository.
type Repository interface {
	GetUser(ctx context.Context, login, pass string) (*auth.User, error)
	AuthUser(ctx context.Context, login, pass string) error
	CreateUser(ctx context.Context, user *auth.User) error
}

package server

import (
	"github.com/imega-teleport/auth/api"
	"github.com/imega-teleport/auth/model"
	"golang.org/x/net/context"
)

// Option is any options for package
type Option func(s *srv)

// WithRepo add repository to server.
func WithRepo(repo model.Repository) Option {
	return func(s *srv) {
		s.repo = repo
	}
}

type srv struct {
	repo model.Repository
}

func (s *srv) Auth(context.Context, *auth.AuthRequest) (*auth.AuthResponse, error) {
	return &auth.AuthResponse{}, nil
}

// NewServer returns server with given options applied.
func NewServer(opts ...Option) auth.AuthBasicServer {
	s := &srv{}
	for _, opt := range opts {
		opt(s)
	}

	return s
}

package server

import (
	"github.com/google/uuid"
	"github.com/imega-teleport/auth/api"
	"github.com/imega-teleport/auth/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *srv) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	user := &auth.User{
		Login:  uuid.New().String(),
		Pass:   uuid.New().String(),
		Active: true,
	}
	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return &auth.CreateUserResponse{}, err
	}
	return &auth.CreateUserResponse{
		User: user,
	}, nil
}

func (s *srv) GetUser(context.Context, *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	return &auth.GetUserResponse{}, nil
}

func (s *srv) Auth(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	err := s.repo.AuthUser(ctx, req.GetLogin(), req.GetPass())
	if err != nil {
		logrus.Errorf("%s", err)
		return &auth.AuthResponse{}, status.Error(codes.PermissionDenied, "Access denied")
	}

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

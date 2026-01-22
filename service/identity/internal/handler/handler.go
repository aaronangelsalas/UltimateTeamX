package handler

import (
	"context"
	"database/sql"

	identityv1 "UltimateTeamX/proto/identity/v1"
	"UltimateTeamX/service/identity/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IdentityServer struct {
	identityv1.UnimplementedIdentityServiceServer
	DB *sql.DB
}

func (s *IdentityServer) Register(ctx context.Context, req *identityv1.RegisterRequest) (*identityv1.RegisterResponse, error) {

	//func RegisterUser(userRepo *repo.UserRepo, email, username, password string) (string, error)
	id, err := service.RegisterUser(s.DB, req.Email, req.Username, req.Password)
	if err != nil {
		switch err {
		case service.ErrInvalidInput:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case service.ErrEmailExists:
			return nil, status.Error(codes.AlreadyExists, "email already registered")
		case service.ErrUsernameExists:
			return nil, status.Error(codes.AlreadyExists, "username already taken")
		default:
			return nil, status.Error(codes.Internal, "failed to register user")
		}
	}

	return &identityv1.RegisterResponse{UserId: id}, nil
}

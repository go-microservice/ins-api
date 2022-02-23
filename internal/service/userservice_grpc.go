package service

import (
	"context"

	userv1 "github.com/go-microservice/user-service/api/micro/user/v1"

	pb "github.com/go-microservice/ins-api/api/user/v1"
)

var (
	_ pb.UserServiceHTTPServer = (*UserServiceServer)(nil)
)

type UserServiceServer struct {
	userRPC userv1.UserServiceClient
}

func NewUserServiceServer(repo userv1.UserServiceClient) *UserServiceServer {
	return &UserServiceServer{
		userRPC: repo,
	}
}

func (s *UserServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	in := &userv1.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	out, err := s.userRPC.Register(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		Id:       out.Id,
		Username: out.Username,
	}, nil
}

func (s *UserServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *UserServiceServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserServiceServer) BatchGetUsers(ctx context.Context, req *pb.BatchGetUsersRequest) (*pb.BatchGetUsersReply, error) {
	return &pb.BatchGetUsersReply{}, nil
}
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserServiceServer) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	return &pb.UpdatePasswordReply{}, nil
}

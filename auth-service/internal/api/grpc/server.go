package grpc

import (
	"coinflow/coinflow-server/auth-service/internal/api/grpc/types"
	"coinflow/coinflow-server/auth-service/internal/usecases"
	pb "coinflow/coinflow-server/gen/auth_service/golang"
	grpcErr "coinflow/coinflow-server/pkg/pkgerrors/grpc"
	"context"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
	usrService usecases.UserService
}

func NewAuthServer(usrService usecases.UserService) *AuthServer {
	return &AuthServer{
		usrService: usrService,
	}
}

func (s *AuthServer) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	tokens, err := s.usrService.Login(r.Login, r.Password)
	if err != nil {
		return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	return &pb.LoginResponse{AccessToken: tokens.Access, RefreshToken: tokens.Refresh}, nil
}

func (s *AuthServer) Refresh(ctx context.Context, r *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	tokens, err := s.usrService.Refresh(r.RefreshToken)
	if err != nil {
		return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	return &pb.RefreshResponse{AccessToken: tokens.Access, RefreshToken: tokens.Refresh}, nil
}

func (s *AuthServer) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	reqObj, err := types.CreateRegisterRequestObject(r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	usrId, err := s.usrService.Register(reqObj.User)
	if err != nil {
	    return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	return &pb.RegisterResponse{UserId: usrId.String()}, nil
}

func (s *AuthServer) GetUserData(ctx context.Context, r *pb.GetUserDataRequest) (*pb.GetUserDataResponse, error) {
	reqObj, err := types.CreateGetUserDataRequestObject(r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	usr, err := s.usrService.GetUserData(reqObj.UsrId)
	if err != nil {
		return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	resp, err := types.CreateGetUserDataResponse(usr)
	if err != nil {
		return nil, grpcErr.CreateResponseStatusError(err)
	}

	return resp, nil
}

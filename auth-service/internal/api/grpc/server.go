package grpc

import (
	"coinflow/coinflow-server/auth-service/config"
	"coinflow/coinflow-server/auth-service/internal/api/grpc/types"
	"coinflow/coinflow-server/auth-service/internal/usecases"
	pb "coinflow/coinflow-server/gen/auth_service/golang"
	pkgGrpc "coinflow/coinflow-server/pkg/grpc"
	grpcErr "coinflow/coinflow-server/pkg/pkgerrors/grpc"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
	usrService usecases.UserService
	cfg config.ServiceConfig
}

func NewAuthServer(usrService usecases.UserService, cfg config.ServiceConfig) *AuthServer {
	return &AuthServer{
		usrService: usrService,
		cfg: cfg,
	}
}

func (s *AuthServer) setAccessCookie(ctx context.Context, token string) {
	grpc.SetHeader(ctx, metadata.Pairs("Set-Cookie", fmt.Sprintf("%s=%s; Path=/; HttpOnly", s.cfg.AuthCookieName, token)))
}

func (s *AuthServer) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	tokens, err := s.usrService.Login(ctx, r.Login, r.Password)
	if err != nil {
		return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	s.setAccessCookie(ctx, tokens.Access)

	return &pb.LoginResponse{AccessToken: tokens.Access, RefreshToken: tokens.Refresh}, nil
}

func (s *AuthServer) Refresh(ctx context.Context, r *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	tokens, err := s.usrService.Refresh(ctx, r.RefreshToken)
	if err != nil {
		return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	s.setAccessCookie(ctx, tokens.Access)

	return &pb.RefreshResponse{AccessToken: tokens.Access, RefreshToken: tokens.Refresh}, nil
}

func (s *AuthServer) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	reqObj, err := types.CreateRegisterRequestObject(r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	usrId, err := s.usrService.Register(ctx, reqObj.User)
	if err != nil {
	    return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	pkgGrpc.SetResponseCode(ctx, s.cfg.HttpCodeHeaderName, 201)

	return &pb.RegisterResponse{UserId: usrId.String()}, nil
}

func (s *AuthServer) GetUserData(ctx context.Context, r *pb.GetUserDataRequest) (*pb.GetUserDataResponse, error) {
	reqObj, err := types.CreateGetUserDataRequestObject(r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	usr, err := s.usrService.GetUserData(ctx, reqObj.UsrId)
	if err != nil {
		return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	resp, err := types.CreateGetUserDataResponse(usr)
	if err != nil {
		return nil, grpcErr.CreateResponseStatusError(err)
	}

	return resp, nil
}

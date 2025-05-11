package auth

import (
	"coinflow/coinflow-server/api-gateway/internal/clients/grpc/auth/types"
	"coinflow/coinflow-server/api-gateway/internal/models"
	pb "coinflow/coinflow-server/gen/auth_service/golang"
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"context"

	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	grpcCli pb.AuthClient
	grpcCfg pkgConfig.GrpcConfig
}

func NewAuthClient(grpcCfg pkgConfig.GrpcConfig) (*AuthClient, error) {
	const op = "NewAuthClient"

	addr := fmt.Sprintf("%s:%s", grpcCfg.Host, grpcCfg.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &AuthClient{
		grpcCli: pb.NewAuthClient(conn),
		grpcCfg: grpcCfg,
	}, nil
}

func (c *AuthClient) Login(login, password string) (string, string, error) {
	const op = "AuthClient.Login"

	req := pb.LoginRequest{Login: login, Password: password}
	resp, err := c.grpcCli.Login(context.Background(), &req)

	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.AccessToken, resp.RefreshToken, nil
}

func (c *AuthClient) Refresh(refreshToken string) (string, string, error) {
	const op = "AuthClient.Refresh"

	req := pb.RefreshRequest{RefreshToken: refreshToken}
	resp, err := c.grpcCli.Refresh(context.Background(), &req)

	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.AccessToken, resp.RefreshToken, nil
}

func (c *AuthClient) Register(usr *models.User) (uuid.UUID, error) {
	const op = "AuthClient.Register"

	req, err := types.CreateRegisterRequest(usr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	resp, err := c.grpcCli.Register(context.Background(), req)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	id, err := uuid.Parse(resp.UserId)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (c *AuthClient) GetUserData(usrId uuid.UUID) (*models.User, error) {
	const op = "AuthClient.GetUserData"

	req := pb.GetUserDataRequest{UserId: usrId.String()}
	resp, err := c.grpcCli.GetUserData(context.Background(), &req)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	usr, err := types.CreateGetUserDataResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return usr, nil
}

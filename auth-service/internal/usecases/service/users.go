package service

import (
	"coinflow/coinflow-server/auth-service/config"
	"coinflow/coinflow-server/auth-service/internal/models"
	"coinflow/coinflow-server/auth-service/internal/repository"
	"coinflow/coinflow-server/auth-service/internal/usecases"
	"coinflow/coinflow-server/pkg/infra/cache"
	pkgCrypto "coinflow/coinflow-server/pkg/utils/crypto"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	usersRepo repository.UsersRepo
	cache cache.Cache
	jwtCfg config.JwtConfig
	jwtKeys *pkgCrypto.JwtKeys
}

func NewUserService(
	usersRepo repository.UsersRepo,
	cache cache.Cache,
	jwtCfg config.JwtConfig,
) (*UserService, error) {
	const op = "NewUserService"

	privateKey, err := base64.StdEncoding.DecodeString(jwtCfg.PrivateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	publicKey, err := base64.StdEncoding.DecodeString(jwtCfg.PublicKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &UserService{
		usersRepo: usersRepo,
		cache: cache,
		jwtCfg: jwtCfg,
		jwtKeys: &pkgCrypto.JwtKeys{PrivateKey: privateKey, PublicKey: publicKey},
	}, nil
}

func (s *UserService) GenerateNewTokenPair(usrId uuid.UUID) (*usecases.TokenPair, error) {
	const op = "UserService.GenerateNewTokenPair"

	access, err := pkgCrypto.GenerateJwtToken(usrId, time.Now().Add(s.jwtCfg.AccessExpirationTime), s.jwtKeys)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	refresh, err := pkgCrypto.GenerateJwtToken(usrId, time.Now().Add(s.jwtCfg.RefreshExpirationTime), s.jwtKeys)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &usecases.TokenPair{Access: access, Refresh: refresh}, nil
}

func (s *UserService) Login(login, password string) (*usecases.TokenPair, error) {
	const op = "UserService.Login"

	usr, err := s.usersRepo.GetUserByCred(login, password)
	if err != nil {
		return nil, err
	}

	tokens, err := s.GenerateNewTokenPair(usr.Id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tokens, nil
}

func (s *UserService) Refresh(refreshToken string) (*usecases.TokenPair, error) {
	const op = "UserService.Refresh"

	_, err := s.cache.Get(context.Background(), refreshToken)

	if err == nil {
		return nil, fmt.Errorf("%s: %w", op, usecases.ErrorTokenInBlacklist)
	} else if !errors.Is(err, cache.ErrorKeyNotFound) {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	
	usrId, err := pkgCrypto.ValidateJwtToken(refreshToken, s.jwtKeys)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.cache.Set(context.Background(), refreshToken, "", s.jwtCfg.RefreshExpirationTime + time.Minute)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	tokens, err := s.GenerateNewTokenPair(usrId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tokens, nil
}

func (s *UserService) Register(usr *models.User) (uuid.UUID, error) {
	return s.usersRepo.PostUser(usr)
}

func (s *UserService) GetUserData(usrId uuid.UUID) (*models.User, error) {
	return s.usersRepo.GetUser(usrId)
}

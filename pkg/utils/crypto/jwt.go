package crypto

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtKeys struct {
	PrivateKey 		[]byte
	PublicKey 		[]byte
}

var (
	method = jwt.SigningMethodEdDSA
)

func ValidateJwtToken(tokenStr string, keys *JwtKeys) (uuid.UUID, error) {
	const op = "ValidateJwtToken"

	keyFunc := func(token *jwt.Token) (any, error) {
		return ed25519.PublicKey(keys.PublicKey), nil
	}

	token, err := jwt.Parse(tokenStr, keyFunc,
		jwt.WithValidMethods([]string{method.Alg()}),
		jwt.WithExpirationRequired(),
	)

	expTime, err := token.Claims.GetExpirationTime()

	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	} else if time.Now().After(expTime.Time) {
		return uuid.Nil, fmt.Errorf("%s: %w", op, ErrorTokenExpired)
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	usrId, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return usrId, nil
}

func GenerateJwtToken(usrId uuid.UUID, expiresAt time.Time, keys *JwtKeys) (string, error) {
	const op = "GenerateJwtToken"

	claims := jwt.RegisteredClaims{
		Subject: usrId.String(),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token := jwt.NewWithClaims(method, claims)
	tokenStr, err := token.SignedString(ed25519.PrivateKey(keys.PrivateKey))

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tokenStr, nil
}

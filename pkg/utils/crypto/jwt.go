package crypto

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	method = jwt.SigningMethodEdDSA
)

func ValidateJwtToken(tokenStr string, publicKey []byte) (uuid.UUID, error) {
	const op = "ValidateJwtToken"

	keyFunc := func(token *jwt.Token) (any, error) {
		if token.Method != method {
			return nil, ErrorUnexpectedSigningMethod
		}

		return ed25519.PublicKey(publicKey), nil
	}

	token, err := jwt.Parse(tokenStr, keyFunc,
		jwt.WithValidMethods([]string{method.Alg()}),
		jwt.WithExpirationRequired(),
	)

	if errors.Is(err, jwt.ErrTokenExpired) {
		return uuid.Nil, fmt.Errorf("%s: %w", op, ErrorTokenExpired)
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return uuid.Nil, fmt.Errorf("%s: %w", op, ErrorTokenSignatureInvalid)
	} else if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, ErrorTokenParsingFailed)
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

func GenerateJwtToken(usrId uuid.UUID, expiresAt time.Time, privateKey []byte) (string, error) {
	const op = "GenerateJwtToken"

	claims := jwt.RegisteredClaims{
		Subject: usrId.String(),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID: uuid.New().String(),
	}

	token := jwt.NewWithClaims(method, claims)
	tokenStr, err := token.SignedString(ed25519.PrivateKey(privateKey))

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tokenStr, nil
}

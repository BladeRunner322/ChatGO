package core_auth

import (
	"errors"
	"fmt"
	"time"

	core_errors "github.com/BladeRunner322/ChatGO/internal/core/errors"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret []byte
	ttl    time.Duration
}

func NewJWT(cfg Config) *JWT {
	return &JWT{
		secret: []byte(cfg.Secret),
		ttl:    time.Duration(cfg.TTL) * time.Hour,
	}
}

func (j *JWT) Generate(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(j.ttl).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(j.secret)
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return signed, nil
}

func (j *JWT) Parse(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, core_errors.ErrInvalidSigningMethod
		}

		return j.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, fmt.Errorf("token expired: %w", core_errors.ErrExpiredToken)
		}

		return 0, fmt.Errorf("parse token: %w", core_errors.ErrInvalidToken)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, fmt.Errorf("missing user_id claim: %w", core_errors.ErrInvalidToken)
		}

		return int(userID), nil
	}

	return 0, fmt.Errorf("invalid token: %w", core_errors.ErrInvalidToken)
}

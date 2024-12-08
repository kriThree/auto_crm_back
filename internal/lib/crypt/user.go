package user_crypt

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID int64
	Role   string
	jwt.RegisteredClaims
}

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, duration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     secretKey,
		tokenDuration: duration,
	}
}

func (j *JWTManager) Generate(_ context.Context, userID int64, role string) (string, string, error) {
	accessClaims := &UserClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	refreshClaims := &UserClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration * 10)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken,err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(j.secretKey))
	if err != nil {
		return "", "", err
	}
	refreshToken,err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(j.secretKey))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (j *JWTManager) Validate(ctx context.Context, tokenStr string) (userId int64, role string, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return 0,"", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*UserClaims)

	if ok && token.Valid {
		return claims.UserID,claims.Role, nil
	} else {
		return 0,"", errors.New("incorrect token")
	}
}

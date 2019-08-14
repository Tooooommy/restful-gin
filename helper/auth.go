package helper

import (
	"CrownDaisy_GOGIN/config"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	AccountID string // 用户账号ID
	SessionID string // 存储ID
	jwt.StandardClaims
}

func GenAccessToken(accountID string) (string, error) {
	cfg := config.Get().Jwt
	claims := CustomClaims{
		accountID,
		UUID(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(cfg.Expired) * time.Hour).Unix(),
			Issuer:    cfg.Issuer,
			Id:        UUID(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return t.SignedString([]byte(cfg.Secret))
}

func ParseAccessToken(accessToken string) (string, string, error) {
	var claims = &CustomClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Get().App.Secret), nil
	})

	if !token.Valid {
		err = errors.New("token invalid")
	}

	return claims.AccountID, claims.SessionID, err
}

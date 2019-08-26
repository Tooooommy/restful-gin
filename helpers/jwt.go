package helpers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"restful-gin/config"
	"time"
)

type CustomClaims struct {
	AccountID string // 用户账号ID
	SessionID string // 存储ID
	jwt.StandardClaims
}

func GenAccessToken(accountId string, sessionId string) (string, error) {
	cfg := config.Get().Jwt
	claims := CustomClaims{
		accountId,
		sessionId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(cfg.Expired) * time.Hour).Unix(),
			Id:        UUID(),
			Issuer:    cfg.Issuer,
			Subject:   cfg.Subject,
			Audience:  cfg.Audience,
			NotBefore: cfg.NotBefore,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return t.SignedString([]byte(cfg.Secret))
}

func ParseAccessToken(accessToken string) (*CustomClaims, error) {
	cfg := config.Get().Jwt
	var claims = &CustomClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Get().App.Secret), nil
	})

	if token == nil || !token.Valid {
		return claims, errors.New("token invalid")
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return claims, errors.New("token expired")
	}
	if !claims.VerifyIssuer(cfg.Issuer, true) {
		return claims, errors.New("token issuer invalid")
	}
	if !claims.VerifyNotBefore(1, true) {
		return claims, errors.New("token not before invalid")
	}
	if !claims.VerifyAudience(cfg.Audience, true) {
		return claims, errors.New("token audience invalid")
	}
	return claims, err
}

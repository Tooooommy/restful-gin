package base_controller

import (
	"CrownDaisy_GOGIN/helpers"
	"github.com/gin-gonic/gin"
	"strings"
)

const AccessTokenType = "Bearer"

func (auth *JwtAuth) MidJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := auth.GetAccessToken(c)
		claims, err := helpers.ParseAccessToken(accessToken)
		helpers.CheckErr(err, helpers.ReturnResult(helpers.CodeAccessTokenInvalid, "access token invalid", nil))
		c.Set("account_id", claims.AccountID)
		c.Set("session_id", claims.SessionID)
		c.Next()
	}
}

func (auth *JwtAuth) GetAccessToken(c *gin.Context) string {
	var accessToken string
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader != "" {
		tokenParts := strings.Split(tokenHeader, " ")
		helpers.Assert(len(tokenParts) == 2, helpers.ReturnResult(helpers.CodeAccessTokenInvalid, "access token too short.", nil))
		helpers.Assert(tokenParts[0] == AccessTokenType, helpers.ReturnResult(helpers.CodeAccessTokenInvalid, "access token is not bear.", nil))
		accessToken = tokenParts[1]
	} else {
		accessToken = c.Query("access_token")
	}
	helpers.Assert(len(accessToken) > 0, helpers.ReturnResult(helpers.CodeAccessTokenInvalid, "access token is empty.", nil))
	return accessToken
}

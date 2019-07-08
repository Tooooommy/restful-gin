package middleware

import (
	"CrownDaisy_GOGIN/helper"
	"github.com/gin-gonic/gin"
	"strings"
)

const AccessTokenType = "Bearer"

func MidJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := GetAccessToken(c)
		accountID, sessionID, err := helper.ParseAccessToken(accessToken)
		helper.CheckErr(err, helper.ReturnResult(helper.CodeAccessTokenInvalid, err.Error(), nil))
		c.Set("account_id", accountID)
		c.Set("session_id", sessionID)
		c.Next()
	}
}

func GetAccessToken(c *gin.Context) string {
	var accessToken string
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader != "" {
		tokenParts := strings.Split(tokenHeader, " ")
		helper.Assert(len(tokenParts) == 2, helper.ReturnResult(helper.CodeAccessTokenInvalid, "access token too short.", nil))
		helper.Assert(tokenParts[0] == AccessTokenType, helper.ReturnResult(helper.CodeAccessTokenInvalid, "access token is not bear.", nil))
		accessToken = tokenParts[1]
	} else {
		accessToken = c.Query("access_token")
	}
	helper.Assert(len(accessToken) > 0, helper.ReturnResult(helper.CodeAccessTokenInvalid, "access token is empty.", nil))
	return accessToken
}

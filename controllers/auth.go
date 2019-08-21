package base_controller

import (
	"CrownDaisy_GOGIN/db/model"
	"CrownDaisy_GOGIN/helpers"
	"CrownDaisy_GOGIN/helpers/define"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type AuthCtl struct {
	*BaseCtl
	*jwt.GinJWTMiddleware
}

type AccountParams struct {
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	SessionID string `json:"-"`
}

type AccountResult struct {
	AccessToken string
	Expire      time.Time
}

func InitAuth() (*AuthCtl, error) {
	jwtMid, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "zero zone",
		SigningAlgorithm: "HS512",
		Key:              []byte("十步杀一人，千里不留行"),
		Timeout:          24 * time.Hour,
		MaxRefresh:       1 * time.Hour,
		Authenticator:    Authenticator,
		Authorizator:     Authorization,
		PayloadFunc:      PayloadFunc,
		Unauthorized:     Unauthorized,
		LoginResponse:    LoginResponse,
		RefreshResponse:  RefreshResponse,
		IdentityHandler:  IdentityHandler,
		IdentityKey:      "id",
		TokenLookup:      "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:    "Bearer",
		TimeFunc:         time.Now,
	})
	return &AuthCtl{new(BaseCtl), jwtMid}, err
}

// login flow
// authenticator
// Payload func
// login response

func Authenticator(c *gin.Context) (interface{}, error) {
	var accountParams AccountParams
	// get account json from web
	// get verify account from database
	// gen uuid session id into redis
	err := c.BindJSON(&accountParams)
	helpers.CheckErr(err, helpers.ReturnResult(define.MissBindValues, "miss bind values", nil))

	// username : name or email
	var password = helpers.GenPwd(accountParams.Password)
	var account = &model.AccountModel{}
	var exist = false
	if strings.Contains(accountParams.UserName, "@") {
		account.Email = accountParams.UserName
	} else {
		account.Name = accountParams.UserName
	}
	account.Password = password
	account, exist = account.GetAccount()
	helpers.Assert(exist, helpers.ReturnResult(define.NotExistAccount, "account not exist", nil))
	// 保存sessionID
	sessionID := helpers.UUID()
	account.SessionId = sessionID
	accountParams.SessionID = sessionID
	accountParams.UserName = account.Name
	return accountParams, nil
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	if account, ok := data.(AccountParams); ok {
		return jwt.MapClaims{
			"account": account.UserName,
			"session": account.SessionID,
		}
	}
	return jwt.MapClaims{}
}

func LoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(code, helpers.ReturnResult(define.Success, "success", &AccountResult{
		token, expire,
	}))
}

//jwt flow
//IdentityHandler
//Authorization
//Unauthorized

func IdentityHandler(c *gin.Context) interface{} {
	// 验证 session 是否一样
	claims := jwt.ExtractClaims(c)
	return &AccountParams{
		UserName:  claims["account"].(string),
		SessionID: claims["session"].(string),
	}
}

func Authorization(data interface{}, c *gin.Context) bool {
	if accountParams, ok := data.(AccountParams); ok {
		// 判断session是否一致
		account := model.AccountModel{Name: accountParams.UserName, SessionId: accountParams.SessionID}
		_, exist := account.GetAccount()
		c.Set("account_id", account.UniqueId)
		c.Set("session_id", account.SessionId)
		return exist
	}
	return false
}

func Unauthorized(c *gin.Context, code int, msg string) {
	c.JSON(code, helpers.ReturnResult(define.AuthAccountInvalid, msg, nil))
}

func RefreshResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(code, helpers.ReturnResult(define.Success, "success", &AccountResult{token, expire}))
}

const AccessTokenType = "Bearer"

func (auth *AuthCtl) MidJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := auth.GetAccessToken(c)
		claims, err := helpers.ParseAccessToken(accessToken)
		helpers.CheckErr(err, helpers.ReturnResult(define.AccessTokenInvalid, "access token invalid", nil))
		c.Set("account_id", claims.AccountID)
		c.Set("session_id", claims.SessionID)
		c.Next()
	}
}

func (auth *AuthCtl) GetAccessToken(c *gin.Context) string {
	var accessToken string
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader != "" {
		tokenParts := strings.Split(tokenHeader, " ")
		helpers.Assert(len(tokenParts) == 2, helpers.ReturnResult(define.AccessTokenInvalid, "access token too short.", nil))
		helpers.Assert(tokenParts[0] == AccessTokenType, helpers.ReturnResult(define.AccessTokenInvalid, "access token is not bear.", nil))
		accessToken = tokenParts[1]
	} else {
		accessToken = c.Query("access_token")
	}
	helpers.Assert(len(accessToken) > 0, helpers.ReturnResult(define.AccessTokenInvalid, "access token is empty.", nil))
	return accessToken
}

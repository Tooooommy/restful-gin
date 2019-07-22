package account_controller

import (
	"CrownDaisy_GOGIN/db/model"
	"CrownDaisy_GOGIN/helper"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type AccountParams struct {
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	SessionID string `json:"-"`
}

type AccountResult struct {
	AccessToken string
	Expire      time.Time
}

var JwtMid *jwt.GinJWTMiddleware

func InitMidAuth() (err error) {
	JwtMid, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "zero zone",
		SigningAlgorithm: "HS512",
		Key:              []byte("十步杀一人，千里不留行"),
		Timeout:          24 * time.Hour,
		MaxRefresh:       1 * time.Hour,
		Authenticator:    Authenticator,
		Authorizator:     Authorizator,
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
	return
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
	helper.CheckErr(err, helper.ReturnResult(helper.CodeMissBindValues, "miss bind values", nil))

	// username : name or email
	var password = helper.GenPwd(accountParams.Password)
	var account = &model.AccountModel{}
	var exist = false
	if strings.Contains(accountParams.UserName, "@") {
		account.Email = accountParams.UserName
	} else {
		account.Name = accountParams.UserName
	}
	account.Password = password
	account, exist = account.GetAccount()
	helper.Assert(exist, helper.ReturnResult(helper.CodeNotExistAccount, "account not exist", nil))
	// 保存sessionID
	sessionID := helper.UUID()
	account.SessionID = sessionID
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
	c.JSON(code, helper.ReturnResult(helper.CodeSuccess, "success", &AccountResult{
		token, expire,
	}))
}

//jwt flow
//IdentityHandler
//Authorizator
//Unauthorized

func IdentityHandler(c *gin.Context) interface{} {
	// 验证 session 是否一样
	claims := jwt.ExtractClaims(c)
	return &AccountParams{
		UserName:  claims["account"].(string),
		SessionID: claims["session"].(string),
	}
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if accountParams, ok := data.(AccountParams); ok {
		// 判断session是否一致
		account := model.AccountModel{Name: accountParams.UserName, SessionID: accountParams.SessionID}
		_, exist := account.GetAccount()
		return exist
	}
	return false
}

func Unauthorized(c *gin.Context, code int, msg string) {
	c.JSON(code, helper.ReturnResult(helper.CodeAuthAccountInvalid, msg, nil))
}

func RefreshResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(code, helper.ReturnResult(helper.CodeSuccess, "success", &AccountResult{token, expire}))
}

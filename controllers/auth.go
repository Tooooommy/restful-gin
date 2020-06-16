package base_ctl

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"restful-gin/db/model"
	"restful-gin/helpers"
	"restful-gin/helpers/define"
	"strings"
	"time"
)

var Auth = initAuth()

type AuthCtl struct {
	*BaseCtl
	*jwt.GinJWTMiddleware
}

type AccountParams struct {
	Username  string `json:"username"` // 用户名字
	Password  string `json:"password"`
	SessionID string `json:"-"`
}

type AccountResult struct {
	AccessToken string
	Expire      time.Time
}

func initAuth() *AuthCtl {
	jwtMid, err := jwt.New(&jwt.GinJWTMiddleware{
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
	if err != nil {
		panic(err)
	}
	return &AuthCtl{new(BaseCtl), jwtMid}
}

// login flow
// authenticator
// Payload func
// login response
func Authenticator(c *gin.Context) (interface{}, error) {
	var params AccountParams
	err := c.BindJSON(&params)
	helpers.CheckErr(err, helpers.ReturnResult(define.MissBindValues, "miss bind values", nil))

	// username : name or email
	var password = helpers.GenPwd(params.Password)
	var account = &model.AccountModel{}
	var exist = false
	if strings.Contains(params.Username, "@") {
		account.Email = params.Username
	} else {
		account.Name = params.Username
	}
	account.Password = password
	account, exist = account.GetAccount()
	helpers.Assert(exist, helpers.ReturnResult(define.NotExistAccount, "account not exist", nil))
	// 保存sessionID
	sessionID := helpers.UUID()
	account.SessionId = sessionID
	params.SessionID = sessionID
	params.Username = account.Name
	return params, nil
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	if account, ok := data.(AccountParams); ok {
		return jwt.MapClaims{
			"account": account.Username,
			"session": account.SessionID,
		}
	}
	return jwt.MapClaims{}
}

func LoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(code, helpers.ReturnResult(define.Success, "success", &AccountResult{token, expire,}))
}

//jwt flow
//IdentityHandler
//Authorization
//Unauthorized
func IdentityHandler(c *gin.Context) interface{} {
	// 验证 session 是否一样
	claims := jwt.ExtractClaims(c)
	return &AccountParams{
		Username:  claims["account"].(string),
		SessionID: claims["session"].(string),
	}
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if params, ok := data.(AccountParams); ok {
		// 判断session是否一致
		account := model.AccountModel{Name: params.Username, SessionId: params.SessionID}
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


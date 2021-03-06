package account_ctl

import (
	"fmt"
	"net/http"
	"restful-gin/config"
	base_controller "restful-gin/controllers"
	"restful-gin/helpers"
	"restful-gin/helpers/define"
	"restful-gin/libs/qq"
	"restful-gin/libs/wechat"
	types "restful-gin/type"

	"github.com/gin-gonic/gin"
)

type AccountCtl struct {
	base_controller.BaseCtl
}

// redirect to wechat
func (t *AccountCtl) RedirectWeChatLoginPage(c *gin.Context) {
	cfg := config.Get().WeChat
	cfg.State = helpers.UUID()
	auth := wechat.New(cfg.AppId, cfg.RedirectUri, cfg.State, cfg.Scope)
	auth.RedirectWithGin(c)
	return
}

func (t *AccountCtl) AuthWeChatCallback(c *gin.Context) {
	// wechat auth redirect uri
	cfg := config.Get().WeChat
	redirectUri := c.Query("redirect_uri")
	if cfg.RedirectUri != redirectUri {
		result := types.ReturnResult(define.AuthRedirectUriInvalid, "auth redirect uri invalid", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 验证state 判断是不是此次的授权
	state := c.Query("state")
	if state == "" {
		result := types.ReturnResult(define.AuthStateInvalid, "auth state invalid", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	code := c.Query("code")
	if code == "" {
		result := types.ReturnResult(define.AuthCodeEmpty, "auth code is empty", nil)
		c.JSON(http.StatusOK, result)
		return
	}

	client := wechat.DefaultClient(cfg.AppId, cfg.AppSecret, code)
	token, err := client.GetAccessToken(code)
	if err != nil {
		result := types.ReturnResult(define.AuthAccessTokenError, "auth access token failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 保存token
	fmt.Printf("%+v", token)
	// 获取user info
	userInfo, err := client.GetUserInfo(token.AccessToken, token.Openid, cfg.Lang)
	if err != nil {
		result := types.ReturnResult(define.AuthUserInfoError, "auth user info failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 保存user info
	fmt.Printf("%+v", userInfo)
	return
}

// redirect to wechat
func (t *AccountCtl) RedirectQQLoginPage(c *gin.Context) {
	cfg := config.Get().QQ
	cfg.State = helpers.UUID()
	auth := qq.New(cfg.ClientId, cfg.ClientSecret, cfg.RedirectUri, cfg.State, cfg.Scope)
	auth.RedirectWithGin(c)
	return
}

func (t *AccountCtl) AuthQQCallback(c *gin.Context) {
	// wechat auth redirect uri
	cfg := config.Get().QQ
	// 验证state 判断是不是此次的授权
	state := c.Query("state")
	if state == "" {
		result := types.ReturnResult(define.AuthStateInvalid, "auth state invalid", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	code := c.Query("code")
	if code == "" {
		result := types.ReturnResult(define.AuthCodeEmpty, "auth code is empty", nil)
		c.JSON(http.StatusOK, result)
		return
	}

	auth := qq.New(cfg.ClientId, cfg.ClientSecret, cfg.RedirectUri, cfg.State, cfg.Display)
	client := qq.DefaultClient(auth)
	token, err := client.GetAccessToken(code)
	if err != nil {
		result := types.ReturnResult(define.AuthAccessTokenError, "auth access token failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 保存token
	fmt.Printf("%+v", token)
	// get openid
	openMe, err := client.GetOpenId(token.AccessToken)
	if err != nil {
		result := types.ReturnResult(define.AuthUserInfoError, "auth openid gain failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 获取user info
	userInfo, err := client.GetUserInfo(token.AccessToken, openMe.Openid)
	if err != nil {
		result := types.ReturnResult(define.AuthUserInfoError, "auth user info failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 保存user info
	fmt.Printf("%+v", userInfo)
	return
}

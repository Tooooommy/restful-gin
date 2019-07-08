package account_controller

import (
	"CrownDaisy_GOGIN/config"
	"CrownDaisy_GOGIN/helper"
	"CrownDaisy_GOGIN/lib/qq"
	"CrownDaisy_GOGIN/lib/wechat"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

// redirect to wechat
func (*AuthController) RedirectWeChatLoginPage(c *gin.Context) {
	cfg := config.Get().WeChat
	cfg.State = helper.UUID()
	auth := wechat.New(cfg.AppId, cfg.RedirectUri, cfg.State, cfg.Scope)
	auth.RedirectWithGin(c)
	return
}

func (*AuthController) AuthWeChatCallback(c *gin.Context) {
	// wechat auth redirect uri
	cfg := config.Get().WeChat
	redirectUri := c.Query("redirect_uri")
	if cfg.RedirectUri != redirectUri {
		result := helper.ReturnResult(helper.CodeAuthRedirectUriInvalid, "auth redirect uri invalid", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 验证state 判断是不是此次的授权
	state := c.Query("state")
	if state == "" {
		result := helper.ReturnResult(helper.CodeAuthStateInvalid, "auth state invalid", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	code := c.Query("code")
	if code == "" {
		result := helper.ReturnResult(helper.CodeAuthCodeEmpty, "auth code is empty", nil)
		c.JSON(http.StatusOK, result)
		return
	}

	client := wechat.DefaultClient(cfg.AppId, cfg.AppSecret, code)
	token, err := client.GetAccessToken(code)
	if err != nil {
		result := helper.ReturnResult(helper.CodeAuthAccessTokenError, "auth access token failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 保存token
	fmt.Printf("%+v", token)
	// 获取user info
	userInfo, err := client.GetUserInfo(token.AccessToken, token.Openid, cfg.Lang)
	if err != nil {
		result := helper.ReturnResult(helper.CodeAuthUserInfoError, "auth user info failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 保存user info
	fmt.Printf("%+v", userInfo)
	return
}

// redirect to wechat
func (*AuthController) RedirectQQLoginPage(c *gin.Context) {
	cfg := config.Get().QQ
	cfg.State = helper.UUID()
	auth := qq.New(cfg.ClientId, cfg.ClientSecret, cfg.RedirectUri, cfg.State, cfg.Scope)
	auth.RedirectWithGin(c)
	return
}

func (*AuthController) AuthQQCallback(c *gin.Context) {
	// wechat auth redirect uri
	cfg := config.Get().QQ
	// 验证state 判断是不是此次的授权
	state := c.Query("state")
	if state == "" {
		result := helper.ReturnResult(helper.CodeAuthStateInvalid, "auth state invalid", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	code := c.Query("code")
	if code == "" {
		result := helper.ReturnResult(helper.CodeAuthCodeEmpty, "auth code is empty", nil)
		c.JSON(http.StatusOK, result)
		return
	}

	auth := qq.New(cfg.ClientId, cfg.ClientSecret, cfg.RedirectUri, cfg.State, cfg.Display)
	client := qq.DefaultClient(auth)
	token, err := client.GetAccessToken(code)
	if err != nil {
		result := helper.ReturnResult(helper.CodeAuthAccessTokenError, "auth access token failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 保存token
	fmt.Printf("%+v", token)
	// get openid
	openMe, err := client.GetOpenId(token.AccessToken)
	if err != nil {
		result := helper.ReturnResult(helper.CodeAuthUserInfoError, "auth openid gain failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 获取user info
	userInfo, err := client.GetUserInfo(token.AccessToken, openMe.Openid)
	if err != nil {
		result := helper.ReturnResult(helper.CodeAuthUserInfoError, "auth user info failed", nil)
		c.JSON(http.StatusOK, result)
		return
	}
	// 保存user info
	fmt.Printf("%+v", userInfo)
	return
}

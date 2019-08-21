package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
)

const WechatApi = "https://open.weixin.qq.com"

type Auth struct {
	AppID        string
	RedirectUri  string
	ResponseType string
	Scope        string
	State        string
}

func New(appId, redirectUri, state string, scope string) *Auth {
	return &Auth{
		AppID:        appId,
		RedirectUri:  redirectUri,
		ResponseType: "code",
		State:        state,
		Scope:        scope,
	}
}

// redirect url
// https://open.weixin.qq.com/connect/qrconnect?appid=APPID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&state=STATE#wechat_redirect
func (auth *Auth) AuthCodeUrl() string {
	var params = make(url.Values)
	params.Set("appid", auth.AppID)
	params.Set("redirect_uri", auth.RedirectUri)
	params.Set("response_type", auth.ResponseType)
	params.Set("state", auth.State)
	params.Set("scope", auth.Scope)
	return fmt.Sprintf("%s/connect/qrconnect?%s#echat_redirect", WechatApi, params.Encode())
}

func (auth *Auth) Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, auth.AuthCodeUrl(), 302)
}

func (auth *Auth) RedirectWithGin(ctx *gin.Context) {
	ctx.Redirect(302, auth.AuthCodeUrl())
}

// get access token
// https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
type Token struct {
	AccessToken  string
	ExpiresIn    int
	RefreshToken string
	Openid       string
	Scope        string
	Unionid      string
}

type AuthToken struct {
	AppID     string
	Secret    string
	Code      string
	GrantType string
}

type Client struct {
	*http.Client
	*AuthToken
}

func DefaultClient(appId, secret, code string) *Client {
	authToken := AuthToken{
		AppID:     appId,
		Secret:    secret,
		Code:      code,
		GrantType: "authorization_code",
	}

	return &Client{http.DefaultClient, &authToken}
}

func (c *Client) readFromBody(body io.ReadCloser, res interface{}) error {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf.Bytes(), res)
	if err != nil {
		return errors.New(fmt.Sprintf("error occurred: %s", buf.String()))
	}
	return nil
}

func (c *Client) GetAccessToken(code string) (*Token, error) {
	c.Code = code
	var params = make(url.Values)
	params.Set("appid", c.AppID)
	params.Set("secret", c.Secret)
	params.Set("code", c.Code)
	params.Set("grant_type", "authorization_code") // 默认
	accessTokenUrl := fmt.Sprintf("%s/sns/oauth2/access_token?%s", WechatApi, params.Encode())
	resp, err := c.Get(accessTokenUrl)
	if err != nil {
		return nil, err
	}
	token := &Token{}
	err = c.readFromBody(resp.Body, token)
	return token, err
}

// refresh token
// https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN
func (c *Client) RefreshAccessToken(refreshToken string) (*Token, error) {
	var params = make(url.Values)
	params.Set("appid", c.AppID)
	params.Set("grant_type", "refresh_token")
	params.Set("refresh_token", refreshToken)
	refreshTokenUrl := fmt.Sprintf("%s/sns/oauth2/refresh_token?%s", WechatApi, params.Encode())
	resp, err := c.Get(refreshTokenUrl)
	if err != nil {
		return nil, err
	}
	token := &Token{}
	err = c.readFromBody(resp.Body, token)
	return token, err
}

// check access token
//https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID
type AccessTokenError struct {
	ErrCode int
	ErrMsg  string
}

func (c *Client) CheckAccessToken(token, openid string) (*AccessTokenError, error) {
	var params = make(url.Values)
	params.Set("access_token", token)
	params.Set("openid", openid)
	checkTokenUrl := fmt.Sprintf("%s/sns/auth?%s", WechatApi, params.Encode())
	resp, err := c.Get(checkTokenUrl)
	if err != nil {
		return nil, err
	}
	var accessTokenError = &AccessTokenError{}
	err = c.readFromBody(resp.Body, accessTokenError)
	return accessTokenError, err
}

// get user info
// https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID
type UserInfo struct {
	Openid     string
	Nickname   string
	Sex        string
	Province   string
	City       string
	Country    string
	Headimgurl string
	Privilege  []string
	Unionid    string
}

func (c *Client) GetUserInfo(token, openid, lang string) (*UserInfo, error) {
	var params = make(url.Values)
	params.Set("access_token", token)
	params.Set("openid", openid)
	params.Set("lang", lang)
	userInfoUrl := fmt.Sprintf("%s/sns/userinfo?%s", WechatApi, params.Encode())
	resp, err := c.Get(userInfoUrl)
	if err != nil {
		return nil, err
	}
	var userInfo = &UserInfo{}
	err = c.readFromBody(resp.Body, userInfo)
	return userInfo, err
}

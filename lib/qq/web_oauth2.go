package qq

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

const QQApi = "https://graph.qq.com"

type Auth struct {
	ResponseType string
	ClientId     string
	ClientSecret string
	RedirectUri  string
	State        string
	Scope        string
	Display      string
	GrantType    string
	Code         string
}

type Client struct {
	*http.Client
	*Auth
}

func DefaultClient(auth *Auth) *Client {
	return &Client{http.DefaultClient, auth}
}

func New(clientId, clientSecret, redirectUri, state, display string) *Auth {
	return &Auth{
		ResponseType: "code",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RedirectUri:  redirectUri,
		State:        state,
		Display:      display,
	}
}

// get code
// https://graph.qq.com/oauth2.0/authorize
func (auth *Auth) AuthCodeUrl() string {
	auth.ResponseType = "code"
	var params = make(url.Values)
	params.Set("response_type", auth.ResponseType)
	params.Set("client_id", auth.ClientId)
	params.Set("redirect_uri", auth.RedirectUri)
	params.Set("state", auth.State)
	params.Set("scope", auth.Scope)
	params.Set("display", auth.Display)
	return fmt.Sprintf("%s/oauth2.0/authorize?%s", QQApi, params.Encode())
}

func (auth *Auth) Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, auth.AuthCodeUrl(), 302)
}

func (auth *Auth) RedirectWithGin(ctx *gin.Context) {
	ctx.Redirect(302, auth.AuthCodeUrl())
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

// get code
// https://graph.qq.com/oauth2.0/token
type Token struct {
	AccessToken  string
	ExpiresIn    int
	RefreshToken string
}

func (c *Client) GetAccessToken(code string) (*Token, error) {
	c.Code = code
	c.GrantType = "authorization_code"
	var params = make(url.Values)
	params.Set("grant_type", c.GrantType)
	params.Set("client_id", c.ClientId)
	params.Set("client_secret", c.ClientSecret)
	params.Set("code", c.Code)
	params.Set("redirect_uri", c.RedirectUri)
	accessTokenUrl := fmt.Sprintf("%s/oauth2.0/token?%s", QQApi, params.Encode())
	resp, err := c.Get(accessTokenUrl)
	if err != nil {
		return nil, err
	}
	var token = &Token{}
	err = c.readFromBody(resp.Body, token)
	return token, err
}

// https://graph.qq.com/oauth2.0/token
func (c *Client) RefreshAccessToken(refreshToken string) (*Token, error) {
	c.GrantType = "refresh_token"
	var params = make(url.Values)
	params.Set("grant_type", c.GrantType)
	params.Set("client_id", c.ClientId)
	params.Set("client_secret", c.ClientSecret)
	params.Set("refresh", refreshToken)
	refreshTokenUrl := fmt.Sprintf("%s/oauth2.0/token?%s", QQApi, params.Encode())
	resp, err := c.Get(refreshTokenUrl)
	if err != nil {
		return nil, err
	}
	var token = &Token{}
	err = c.readFromBody(resp.Body, token)
	return token, err
}

// get openid
// https://graph.qq.com/oauth2.0/me
type OpenMe struct {
	ClientId string
	Openid   string
}

func (c *Client) GetOpenId(accessToken string) (*OpenMe, error) {
	var params = make(url.Values)
	params.Set("access_token", accessToken)
	openIdUrl := fmt.Sprintf("%s/oauth2.0/me?%s", QQApi, params.Encode())
	resp, err := c.Get(openIdUrl)
	if err != nil {
		return nil, err
	}
	var openMe = &OpenMe{}
	err = c.readFromBody(resp.Body, openMe)
	return openMe, err
}

// user info
// https://graph.qq.com/user/get_user_info
type UserInfo struct {
	Ret          int
	Msg          string
	Nickname     string
	Figureurl    string
	Figureurl1   string
	Figureurl2   string
	Figureurl3   string
	FigureurlQQ1 string
	FigureurlQQ2 string
	Gender       string
}

func (c *Client) GetUserInfo(token, openid string) (*UserInfo, error) {
	var params = make(url.Values)
	params.Set("access_token", token)
	params.Set("oauth_consumer_key", c.ClientId)
	params.Set("openid", openid)
	userInfoUrl := fmt.Sprintf("%s/user/get_user_info?%s", QQApi, params.Encode())
	resp, err := c.Get(userInfoUrl)
	if err != nil {
		return nil, err
	}
	var userInfo = &UserInfo{}
	err = c.readFromBody(resp.Body, userInfo)
	return userInfo, err
}

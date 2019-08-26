package account_service

import (
	"restful-gin/config"
	"restful-gin/helpers"
	"restful-gin/libs/wechat"
)

func RedirectWeChatLoginPage() string {
	cfg := config.Get().WeChat
	cfg.State = helpers.UUID()
	auth := wechat.New(cfg.AppId, cfg.RedirectUri, cfg.State, cfg.Scope)
	return auth.AuthCodeUrl()
}

package account_service

import (
	"CrownDaisy_GOGIN/config"
	"CrownDaisy_GOGIN/helper"
	"CrownDaisy_GOGIN/lib/wechat"
)

func RedirectWeChatLoginPage() string {
	cfg := config.Get().WeChat
	cfg.State = helper.UUID()
	auth := wechat.New(cfg.AppId, cfg.RedirectUri, cfg.State, cfg.Scope)
	return auth.AuthCodeUrl()
}

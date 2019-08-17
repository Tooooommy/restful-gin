package account_service

import (
	"CrownDaisy_GOGIN/config"
	"CrownDaisy_GOGIN/helpers"
	"CrownDaisy_GOGIN/lib/wechat"
)

func RedirectWeChatLoginPage() string {
	cfg := config.Get().WeChat
	cfg.State = helpers.UUID()
	auth := wechat.New(cfg.AppId, cfg.RedirectUri, cfg.State, cfg.Scope)
	return auth.AuthCodeUrl()
}

package main

import (
	"CrownDaisy_GOGIN/config"
	"CrownDaisy_GOGIN/controllers"
	"CrownDaisy_GOGIN/controllers/account_controller"
	"CrownDaisy_GOGIN/db"
	"CrownDaisy_GOGIN/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	var err error
	// init mysql

	err = db.InitMysqlDB()
	if err != nil {
		logger.Sugar.Errorf("init mysql error: %v", err)
		panic(err)
	}

	err = db.InitRedisDB()
	if err != nil {
		logger.Sugar.Errorf("init redis error: %v", err)
		panic(err)
	}

	// init config
	if err = config.InitConfig(); err != nil {
		logger.Sugar.Panicf("init config error: %v", err)
		panic(err)
	}
	// init jwt auth
	auth, err = base_controller.InitAuth()
	if err != nil {
		logger.Sugar.Errorf("init auth error: %v", err)
		panic(err)
	}

}

var (
	base    *base_controller.BaseCtl
	auth    *base_controller.AuthCtl
	account *account_controller.AccountCtl
)

func addRoutes(router *gin.Engine) {
	router.Use(base.MidCors())
	router.NoRoute(base.NotFound)

	apiRouter := router.Group("/api", base.MidRecovery)
	// no jwt auth
	{
		// login and refresh_token
		apiRouter.POST("/account/login", auth.LoginHandler)
		apiRouter.GET("/account/refresh", auth.RefreshHandler)
		// auth wechat and qq
		apiRouter.GET("/account/auth/wechat", account.RedirectWeChatLoginPage)
		apiRouter.GET("/account/auth/qqconnect", account.RedirectQQLoginPage)
		//auth redirect
		apiRouter.GET("/account/auth/wechat/callback", account.AuthWeChatCallback)
		apiRouter.GET("/account/auth/qq/callback", account.AuthQQCallback)

		// recommend book
	}
	// jwt auth
	authRouter := apiRouter.Group("/auth", auth.MiddlewareFunc())
	{
		authRouter.POST("/logout")
	}

	if err := http.ListenAndServe(config.Get().App.Port, router); err != nil {
		logger.Sugar.Debugf("http router listen error: %+v", err)
		panic(err)
	}
	return
}

func main() {
	addRoutes(gin.Default())
}

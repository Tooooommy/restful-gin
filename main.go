package main

import (
	"CrownDaisy_GOGIN/config"
	"CrownDaisy_GOGIN/controller"
	"CrownDaisy_GOGIN/controller/account_controller"
	"CrownDaisy_GOGIN/db"
	"CrownDaisy_GOGIN/logger"
	"CrownDaisy_GOGIN/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func run() {
	// init config
	config.DefaultConfigPath = "app.ini"
	err := config.InitConf()
	if err != nil {
		fmt.Printf("init config error: %v", err)
	}
	// init logger
	logger.InitLogger()

	// init mysql
	err = db.InitMysqlDB()
	if err != nil {
		fmt.Printf("init mysql error: %v", err)
	}

	// init redis
	err = db.InitRedisDB()
	if err != nil {
		fmt.Printf("init redis error: %v", err)
	}

	// init jwt auth
	err = account_controller.InitMidAuth()
	if err != nil {
		fmt.Printf("init jwt auth error: %+v", err)
	}
}

var (
	auth account_controller.AuthController
)

func addRoutes(router *gin.Engine) {
	jwtMid := account_controller.JwtMid
	router.Use(middleware.MidCors())
	router.NoRoute(controller.NotFound)
	//router.Use(account_controller.JwtMid.MiddlewareFunc())

	apiRouter := router.Group("/api", controller.HandleResultPanic)
	// no jwt auth
	{
		// login and refresh_token
		apiRouter.POST("/account/login", jwtMid.LoginHandler)
		apiRouter.GET("/account/refresh", jwtMid.RefreshHandler)
		// auth wechat and qq
		apiRouter.GET("/account/auth/wechat", auth.RedirectWeChatLoginPage)
		apiRouter.GET("/account/auth/qqconnect", auth.RedirectQQLoginPage)
		//auth redirect
		apiRouter.GET("/account/auth/wechat/callback", auth.AuthWeChatCallback)
		apiRouter.GET("/account/auth/qq/callback", auth.AuthQQCallback)

		// recommend book
	}
	// jwt auth
	authRouter := apiRouter.Group("/auth", jwtMid.MiddlewareFunc())
	{
		authRouter.POST("/logout")
	}

	if err := http.ListenAndServe(":5000", router); err != nil {
		logger.Debugf("http router listen error: %+v", err)
	}
	return
}

func main() {
	run()
	addRoutes(gin.Default())
}

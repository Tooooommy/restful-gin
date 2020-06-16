package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"restful-gin/config"
	"restful-gin/controllers"
	"restful-gin/controllers/account_ctl"
	_ "restful-gin/docs"
	"restful-gin/logger"
)

func init() {
	var err error
	// init mysql

	//err = db.InitMysqlDB()
	//if err != nil {
	//	logger.Sugar.Errorf("init mysql error: %v", err)
	//	panic(err)
	//}
	//
	//err = db.InitRedisDB()
	//if err != nil {
	//	logger.Sugar.Errorf("init redis error: %v", err)
	//	panic(err)
	//}

	// init config
	if err = config.InitConfig(); err != nil {
		logger.Sugar.Panicf("init config error: %v", err)
		panic(err)
	}
}

var (
	base    *base_ctl.BaseCtl
	account *account_ctl.AccountCtl
)
var auth = base_ctl.Auth

func addRoutes(router *gin.Engine) {
	v1 := router.Group("/v1", base.MidRecovery())
	{
		// login and refresh_token
		v1.POST("/account/login", auth.LoginHandler)
		v1.GET("/account/refresh", auth.RefreshHandler)
		// auth wechat and qq
		//apiRouter.GET("/account/auth/wechat", account.RedirectWeChatLoginPage)
		//apiRouter.GET("/account/auth/qqconnect", account.RedirectQQLoginPage)
		//auth redirect
		//apiRouter.GET("/account/auth/wechat/callback", account.AuthWeChatCallback)
		//apiRouter.GET("/account/auth/qq/callback", account.AuthQQCallback)
	}
	authRouter := v1.Group("/auth", auth.MiddlewareFunc())
	{
		authRouter.POST("/logout")
	}

	return
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	router := gin.Default()
	// set swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.NoRoute(base.NotFound)
	router.NoMethod(base.NotFound)
	router.Use(base.MidCors())
	addRoutes(router)
	if err := http.ListenAndServe(config.Get().App.Port, router); err != nil {
		logger.Sugar.Debugf("http router listen error: %+v", err)
		panic(err)
	}
}

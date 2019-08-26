package main

import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"restful-gin/config"
	"restful-gin/controllers"
	"restful-gin/controllers/account_controller"
	"restful-gin/logger"

	"github.com/gin-gonic/gin"
	_ "restful-gin/docs"
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
	// set swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 设置跨域
	router.Use(base.MidCors())

	// 设置无人区404
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
	addRoutes(gin.Default())
}

package base_controller

import (
	"CrownDaisy_GOGIN/define"
	"CrownDaisy_GOGIN/helpers"
	"CrownDaisy_GOGIN/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"time"
)

type BaseCtl struct {
}

func (b *BaseCtl) NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, helpers.ReturnResult(define.NotFoundPage, "page not found", nil))
	return
}

func (b *BaseCtl) MidRecovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.Sugar.Debugf("panic occurred: %+v", r)
			logger.Sugar.Debug(debug.Stack())
			if _, ok := r.(*helpers.Result); ok {
				c.JSON(http.StatusOK, r)
				return
			}
		}
	}()
	c.Next()
}

func (b *BaseCtl) MidCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowFiles:             false,
		AllowBrowserExtensions: false,
		AllowMethods: []string{http.MethodGet, http.MethodPost,
			http.MethodPut, http.MethodDelete, http.MethodOptions,
			http.MethodHead, http.MethodTrace, http.MethodPatch},
		MaxAge: 50 * time.Second,
	})
}

func (b *BaseCtl) Assert(bo bool, res *helpers.Result) {
	helpers.Assert(bo, res)
}

func (b *BaseCtl) CheckErr(err error, res *helpers.Result) {
	helpers.CheckErr(err, res)
}

func (b *BaseCtl) GetAccountId(c *gin.Context) string {
	return c.GetString("account_id")
}

func (b *BaseCtl) GetSessionId(c *gin.Context) string {
	return c.GetString("session_id")
}

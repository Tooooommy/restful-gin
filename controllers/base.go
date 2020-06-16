package base_ctl

import (
	"net/http"
	"restful-gin/helpers/define"
	"restful-gin/logger"
	types "restful-gin/type"
	"runtime/debug"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type BaseCtl struct {
}

func (b *BaseCtl) NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, types.ReturnResult(define.NotFoundPage, "page not found", nil))
	return
}

func (b *BaseCtl) MidRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Sugar.Debugf("panic occurred: %+v", r)
				logger.Sugar.Debug(debug.Stack())
				if _, ok := r.(*types.Result); ok {
					c.JSON(http.StatusOK, r)
					return
				}
			}
		}()
		c.Next()
	}
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

func (b *BaseCtl) Assert(bo bool, code int, message string, ds ...interface{}) {
	types.Assert(bo, code, message, ds)
}

func (b *BaseCtl) CheckErr(err error, cs ...int) {
	types.CheckErr(err, cs...)
}

func (b *BaseCtl) GetAccountId(c *gin.Context) string {
	return c.GetString("account_id")
}

func (b *BaseCtl) GetSessionId(c *gin.Context) string {
	return c.GetString("session_id")
}

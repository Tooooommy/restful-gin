package base_controller

import (
	"CrownDaisy_GOGIN/helper"
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
	c.JSON(http.StatusNotFound, helpers.ReturnResult(helpers.CodeNotFoundPage, "page not found", nil))
	return
}

func (b *BaseCtl) MidRecovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.Sugar.Debugf("panic occurred: %+v", r)
			logger.Sugar.Debug(debug.Stack())
			if _, ok := r.(*helpers.Result); ok {
				c.JSON(http.StatusOK, r)
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

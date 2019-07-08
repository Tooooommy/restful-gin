package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MidCors() gin.HandlerFunc {
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

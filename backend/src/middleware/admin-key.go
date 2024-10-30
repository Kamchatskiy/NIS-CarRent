package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminKeyRequired(adminKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader("admin-key")
		if !strings.EqualFold(key, adminKey) {
			ctx.String(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

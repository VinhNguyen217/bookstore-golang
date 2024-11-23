package middlewares

import (
	"book-store/log"
	"book-store/utils"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"net/http"
)

func Authorization(di *do.Injector) gin.HandlerFunc {
	enforcer := do.MustInvoke[*utils.Enforcer](di)

	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		method := ctx.Request.Method
		path := ctx.Request.URL.Path

		result, _ := enforcer.E.Enforce(role, path, method)
		if result {
			ctx.Next()
		} else {
			log.Warnw(ctx, "casbin check failed", "role", role, "path", path, "method", method)
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}

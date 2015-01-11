// +build !debug

package webui

import "github.com/gin-gonic/gin"

func init() {
	gin.SetMode("release")
}

func TemplateReloader(ctx *gin.Context) {
	ctx.Next()
}

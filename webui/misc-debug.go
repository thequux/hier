// +build debug

package webui

import "github.com/gin-gonic/gin"

func TemplateReloader(ctx *gin.Context) {
	ctx.Engine.HTMLRender.(*TemplateRenderer).Template = LoadTemplates()
	ctx.Next()
}

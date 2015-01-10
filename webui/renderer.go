package webui

import (
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
)

type TemplateRenderer struct {
	Template *template.Template
}

type SkeletonParams struct {
	Title string
	Content interface{}
}

// Concept check
var _ render.Render = &TemplateRenderer{}


func (t *TemplateRenderer) Render(w http.ResponseWriter, code int, data ...interface{}) error {
	file := data[0].(string)
	param := data[1].(SkeletonParams)
	
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	
	err := t.Template.ExecuteTemplate(w, "templates/header.html", param)
	if err == nil {
		err = t.Template.ExecuteTemplate(w, file, param.Content)
	}
	if err == nil {
		err = t.Template.ExecuteTemplate(w, "templates/footer.html", param)
	}
	return err
	
}

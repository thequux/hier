package webui

import (
	"fmt"
	"strings"
	"html/template"
)

func LoadTemplates() *template.Template {
	var tmpl *template.Template = template.New("")
	for _,name := range AssetNames() {
		if strings.HasPrefix(name, "templates/") {
			asset, err := Asset(name)
			if err != nil {
				panic(fmt.Sprintf("Failed to load template %s: %s", name, err))
			}
			_, err = tmpl.New(name).Parse(strings.Replace(string(asset), "\\\n", "", -1))
			if err != nil {
				panic("Failed to parse template " + name + ": " + err.Error())
			}
		}
	}
	return tmpl
}

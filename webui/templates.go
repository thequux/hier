package webui

import (
	"encoding/json"
	"fmt"
	"strings"
	"html/template"
)

var funcMap = map[string]interface{} {
	"json": func(v interface{}) (string, error) {
		out, err := json.Marshal(v)
		return string(out), err
	},
}

func LoadTemplates() *template.Template {
	var tmpl *template.Template = template.New("").Funcs(funcMap)
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

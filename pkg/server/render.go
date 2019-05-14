package server

import (
	"log"
	"os"
	"text/template"
)

// RenderTemplate RenderTemplate
func RenderTemplate(model, tplFilename, swaggerFilename string) {
	f, err := os.OpenFile(swaggerFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0744)
	if err != nil {
		log.Fatal(err)
	}
	tpl := template.Must(template.ParseFiles(tplFilename))
	m := map[string]string{
		"MODEL_NAME": model,
	}
	tpl.Execute(f, m)
}

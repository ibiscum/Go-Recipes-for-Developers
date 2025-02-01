package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.New("tmpl").Parse(`{{range . -}}
{{ if gt . 1 }}
  {{- . }}
{{end -}}
{{end -}}
`)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, []int{-1, 0, 1, 2, 3, 4, 5})
}

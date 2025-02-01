package main

import (
	"html/template"
	"os"
)

const layout = `<!doctype html>
<html lang="en">
  <head>
  <title>{{template "pageTitle" .}}</title>
  </head>
  <body>
  {{template "pageHeader" .}}
  {{template "pageBody" .}}
  {{template "pageFooter" .}}
  </body>
</html>
{{define "pageTitle"}}{{end}}
{{define "pageHeader"}}{{end}}
{{define "pageBody"}}{{end}}
{{define "pageFooter"}}{{end}}
`

const mainPage = `{{define "pageTitle"}}Main Page{{end}}

{{define "pageHeader"}}
<h1>Main page</h1>
{{end}}

{{define "pageBody"}}
This is the page body.
{{end}}

{{define "pageFooter"}}
This is the page footer.
{{end}}
`

const secondPage = `{{define "pageTitle"}}Second page{{end}}

{{define "pageHeader"}}
<h1>Second page</h1>
{{end}}

{{define "pageBody"}}
This is the page body for the second page.
{{end}}`

func main() {
	mainPageTmpl := template.Must(template.New("body").Parse(layout))
	template.Must(mainPageTmpl.Parse(mainPage))

	secondPageTmpl := template.Must(template.New("body").Parse(layout))
	template.Must(secondPageTmpl.Parse(secondPage))
	mainPageTmpl.Execute(os.Stdout, nil)
	secondPageTmpl.Execute(os.Stdout, nil)
}

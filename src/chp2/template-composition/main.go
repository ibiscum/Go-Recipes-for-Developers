package main

import (
	"os"
	"text/template"
)

const tp = `{{define "line"}}
{{.Title}} {{.Author}} {{.PubYear}}
{{end}}
Book list:
{{range . -}}
  {{template "line" .}}
{{end -}}
`

type Book struct {
	Title   string
	Author  string
	PubYear int
}

var books = []Book{
	{
		Title:   "Pride and Prejudice",
		Author:  "Jane Austen",
		PubYear: 1813,
	},
	{
		Title:   "To Kill a Mockingbird",
		Author:  "Harper Lee",
		PubYear: 1960,
	},
	{
		Title:   "The Great Gatsby",
		Author:  "F. Scott Fitzgerald",
		PubYear: 1925,
	},
	{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		PubYear: 1954,
	},
}

func main() {
	tmpl, err := template.New("body").Parse(tp)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, books)
}

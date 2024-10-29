package main

import (
	"os"
	"text/template"
)

type Book struct {
	Title   string
	Author  string
	PubYear int
}

const tp = `The book "{{.Title}}" by {{.Author}} was published in {{.PubYear}}.
`

const tpIter = `{{range .}}
The book "{{.Title}}" by {{.Author}} was published in {{.PubYear}}.
{{end}}`

func main() {
	book1 := Book{
		Title:   "Pride and Prejudice",
		Author:  "Jane Austen",
		PubYear: 1813,
	}
	book2 := Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		PubYear: 1954,
	}
	tmpl, err := template.New("book").Parse(tp)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, book1)
	tmpl.Execute(os.Stdout, book2)

	// Iteration
	tmpl, err = template.New("bookIter").Parse(tpIter)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, []Book{book1, book2})

	tmpl.Execute(os.Stdout, map[int]Book{
		1: book1,
		2: book2})

}

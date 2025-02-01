package main

import (
	"os"
	"text/template"
)

type Book struct {
	Title    string
	Author   string
	Editions []Edition
}

type Edition struct {
	Edition int
	PubYear int
}

const tp = `{{range $bookIndex, $book := .}}
{{$book.Author}}
{{range $book.Editions}}
  {{$book.Title}} Edition: {{.Edition}} {{.PubYear}}
{{end}}
{{end}}`

func main() {
	book1 := Book{
		Title:  "Pride and Prejudice",
		Author: "Jane Austen",
		Editions: []Edition{
			{
				Edition: 1,
				PubYear: 1813,
			},
			{
				Edition: 2,
				PubYear: 1813,
			},
			{
				Edition: 3,
				PubYear: 1817,
			},
		},
	}
	book2 := Book{
		Title:  "The Lord of the Rings",
		Author: "J.R.R. Tolkien",
		Editions: []Edition{
			{
				Edition: 1,
				PubYear: 1954,
			},
			{
				Edition: 2,
				PubYear: 1966,
			},
			{
				Edition: 3,
				PubYear: 1979,
			},
		},
	}
	tmpl, err := template.New("book").Parse(tp)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, []Book{book1, book2})

}

package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"

	"github.com/PacktPublishing/GO-Cookbook-Top-Techniques/src/chp1/internal/routes"
	"github.com/PacktPublishing/GO-Cookbook-Top-Techniques/src/chp1/pkg/commentdb"
)

func main() {
	db, err := sql.Open("sqlite", "webform.db")
	if err != nil {
		panic(err)
	}
	commentdb.InitDB(db)

	r := mux.NewRouter()
	routes.Build(r, db)

	server := http.Server{
		Addr:    ":8181",
		Handler: r,
	}
	server.ListenAndServe()
}

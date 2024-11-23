package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func Build(router *mux.Router, conn *sql.DB) {
	router.Path("/form").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/static/form.html")
	})

	router.Path("/form").Methods("POST").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlePost(conn, w, r)
	})
}

func handlePost(conn *sql.DB, w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	comment := r.PostFormValue("comment")
	_, err := conn.ExecContext(r.Context(), "insert into comments (email,comment) values (?,?)", email, comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/form", http.StatusFound)
}

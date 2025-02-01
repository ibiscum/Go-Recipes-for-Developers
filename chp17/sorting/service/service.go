package service

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	sorting "github.com/ibiscum/1Go-Recipes-for-Developers/chp17/sorting/sort"
)

func HandleSort(w http.ResponseWriter, req *http.Request, ascending bool) {
	var input []time.Time
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(data, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output := sorting.SortTimes(input, ascending)
	data, err = json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /sort/asc", func(w http.ResponseWriter, req *http.Request) {
		HandleSort(w, req, true)
	})
	mux.HandleFunc("POST /sort/desc", func(w http.ResponseWriter, req *http.Request) {
		HandleSort(w, req, false)
	})
	return mux
}

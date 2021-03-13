package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetAppRouter() *mux.Router {
	r := mux.NewRouter()

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("app/static/")))
	r.PathPrefix("/static/").Handler(fs)

	r.HandleFunc("/app", AppHandler).Methods("GET")
	r.HandleFunc("/statistics", StatisticsHandler).Methods("GET")
	r.HandleFunc("/shorten", ShortenHandler)
	r.HandleFunc("/{id}", RedirectHandler).Methods("GET")
	r.HandleFunc("/", IndexHandler).Methods("GET")

	return r
}

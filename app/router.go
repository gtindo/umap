package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetAppRouter() *mux.Router {
	r := mux.NewRouter()

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("app/static/")))
	r.PathPrefix("/static/").Handler(fs)

	r.HandleFunc("/app", AppHandler)
	r.HandleFunc("/statistics", StatisticsHandler)
	r.HandleFunc("/", IndexHandler)

	return r
}

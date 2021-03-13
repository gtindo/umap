package app

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func AppHandler(w http.ResponseWriter, r *http.Request) {
	tmplFiles := []string{
		"app/templates/app.html",
		"app/templates/base.html",
	}

	tmpl := template.Must(template.ParseFiles(tmplFiles...))

	tmpl.Execute(w, nil)
}

func StatisticsHandler(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"app/templates/statistics.html",
		"app/templates/base.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	tmpl.Execute(w, nil)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"app/templates/shorten.html",
		"app/templates/base.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	tmpl.Execute(w, nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/app", 301)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Get Id of url
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, id)
	// Get Id from redis

	// If not exists Get Id from MySQL

	// If exist set this Id on Redis with a timeout

	// Redirect to the url

	// If Id does not exist on MySQL redirect to HTTP 404 page
}

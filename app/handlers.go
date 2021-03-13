package app

import (
	"html/template"
	"net/http"
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Get the Id from URL
	// Get request metadata

	// Check if the Id exists on redis, get url, redirect

	// If not check if it exists on DB, get url, save on redis, redirect

	// If not redirect to http 404 page
}

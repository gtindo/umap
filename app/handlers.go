package app

import (
	"html/template"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
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
	linkId := vars["id"]

	if linkId == "" {
		http.Redirect(w, r, "/app", 301)
	}

	// Get Link from redis
	link, err := rdb.Get(ctx, linkId).Result()

	if err != redis.Nil {
		http.Redirect(w, r, link, 301)
	}

	// If not exists Get Id from MySQL
	_url, err := GetUrl(linkId)

	if err != nil {
		// If Id does not exist on MySQL redirect to HTTP 404 page
		http.Redirect(w, r, "/404", 301)
	}

	// If exist set this Id on Redis with a timeout
	// Redirect to the url
	rdb.Set(ctx, linkId, _url.LinkId, time.Duration(1000))
	http.Redirect(w, r, _url.Link, 301)
}

func Page404Handler(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"app/templates/404.html",
		"app/templates/base.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	tmpl.Execute(w, nil)
}

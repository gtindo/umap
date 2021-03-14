package app

import (
	"html/template"
	"net/http"
	"os"
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

func AnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"app/templates/analytics.html",
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

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	baseUrl := os.Getenv("BASE_URL")

	type ResponseData struct {
		InitialLink string
		NewLink     string
		Error       string
	}

	link := r.FormValue("url")

	data := ResponseData{
		InitialLink: link,
		NewLink:     "",
		Error:       "",
	}

	if !ValidateURI(link) {
		data.Error = "Invalid URL provided"
		tmpl.Execute(w, nil)
		return
	}

	newUrl, err := CreateUrl(link)

	data.NewLink = baseUrl + newUrl.LinkId

	if err != nil {
		data.Error = "Something went wrong !!"
	}

	tmpl.Execute(w, data)
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

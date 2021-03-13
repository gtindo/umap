package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gtindo/umap/app"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	dbStr := os.Getenv("MYSQL_STRING")

	app.InitDb(dbStr)

	defer http.ListenAndServe(":"+port, app.GetAppRouter())

	log.Println("Server listening on port ", port)
}

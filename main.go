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

	redisAddr := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDb := 0

	app.InitDb(dbStr)
	app.InidRedisDb(redisAddr, redisPassword, redisDb)

	defer http.ListenAndServe(":"+port, app.GetAppRouter())

	log.Println("Http Server listening on port", port)
}

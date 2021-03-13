package app

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
)

var db *sql.DB

var ctx = context.Background()
var rdb *redis.Client

func InitDb(dbStr string) {
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connection established with MYSQL DB")
}

func InidRedisDb(addr string, password string, dbNum int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbNum,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	_, err = rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection established with Redis Server")
}

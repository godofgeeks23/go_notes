package main

import (
	"backend/handlers"
	"backend/middlewares"

	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var Port string = ":3000"

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	logger := middlewares.RequestLogger(rdb)
	cache := middlewares.CacheMiddleware(rdb)

	http.HandleFunc("/", logger(cache(handlers.RootHandler)))

	fmt.Println("Server starting at " + Port)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}

}

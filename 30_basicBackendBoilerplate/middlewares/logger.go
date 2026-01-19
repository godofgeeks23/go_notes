package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func RequestLogger(rdb *redis.Client) func(http.HandlerFunc) http.HandlerFunc {

	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next(w, r)

			fmt.Printf("%s %s %s %s\n",
				r.Method,
				r.RequestURI,
				r.RemoteAddr,
				time.Since(start))
		}
	}
}

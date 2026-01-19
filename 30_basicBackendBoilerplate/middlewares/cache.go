package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

const duration = 5 * time.Minute

type responseRecorder struct {
	http.ResponseWriter
	body   []byte
	status int
}

func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.body = append(rec.body, b...)
	return rec.ResponseWriter.Write(b)
}

func (rec *responseRecorder) WriteHeader(statusCode int) {
	rec.status = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

var ctx = context.Background()

func CacheMiddleware(rdb *redis.Client) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				next(w, r)
				return
			}

			key := "cache:" + r.RequestURI

			val, err := rdb.Get(ctx, key).Result()
			if err == nil {
				fmt.Println("cache hit")
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("X-Cache", "HIT")
				w.Write([]byte(val))
				return
			}

			rec := &responseRecorder{ResponseWriter: w, body: []byte{}}
			next(rec, r)

			if rec.status == http.StatusOK {
				rdb.Set(ctx, key, rec.body, duration)
			}
		}
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

func main() {
	router := mux.NewRouter()

	limiter := rate.NewLimiter(rate.Every(time.Second/10), 20)

	router.HandleFunc("/hello", rateLimitMiddleware(limiter, helloHandler)).Methods("GET")

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func rateLimitMiddleware(limiter *rate.Limiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next(w, r)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

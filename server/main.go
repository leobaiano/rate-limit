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

	// Criando um rate limiter com limite de 10 requisições por segundo e capacidade máxima de 20 requisições
	limiter := rate.NewLimiter(rate.Every(time.Second/10), 20)

	// Middleware para aplicar o rate limit à rota "/hello"
	router.HandleFunc("/hello", rateLimitMiddleware(limiter, helloHandler)).Methods("GET")

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Middleware para controle de rate limit
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

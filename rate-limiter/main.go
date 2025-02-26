package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"golang.org/x/time/rate"
)

type Response struct {
	Message string `json:"message"`
}

func getIp(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("Error getting IP address: %v", err)
	}
	return host
}

func rateLimitMiddleware(next http.Handler, limit rate.Limit, burst int) http.Handler {
	ipLimiterMap := make(map[string]*rate.Limiter)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIp(r)

		limiter, ok := ipLimiterMap[ip]
		if !ok {
			limiter = rate.NewLimiter(limit, burst)
			ipLimiterMap[ip] = limiter
		}

		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)

	})
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greetHandler)

	handler := rateLimitMiddleware(mux, rate.Limit(2), 10)
	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}

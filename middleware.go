// middleware.go
package main

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs each incoming request
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
        next.ServeHTTP(w, r)
    })
}

// SetHeadersMiddleware sets common headers for each response
func SetHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

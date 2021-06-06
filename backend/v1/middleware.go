package v1

import (
	"log"
	"net/http"
)

func CorsMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			log.Printf("Option request for %s from %s", r.RequestURI, r.RemoteAddr)
			w.WriteHeader(http.StatusOK)
			return
		}
		// call endpoint
		next.ServeHTTP(w, r)
	})
}

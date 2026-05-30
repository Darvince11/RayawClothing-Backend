package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.RawPath
		log.Println("Url :", data)

		next.ServeHTTP(w, r)
	})
}

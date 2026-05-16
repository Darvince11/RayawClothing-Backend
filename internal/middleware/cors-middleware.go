package middleware

import "net/http"

var AllowedOriging = []string{
	"https://rayaw-store.netlify.app",
	"http://localhost:5173",
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != AllowedOriging[0] && origin != AllowedOriging[1] {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return

		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

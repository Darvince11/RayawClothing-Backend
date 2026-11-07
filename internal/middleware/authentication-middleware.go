package middleware

import (
	"fmt"
	"net/http"
	"rayaw-api/internal/interfaces"
	"slices"
	"strings"
)

func AuthenticationMiddleware(next http.Handler, authInterface interfaces.AuthInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := []string{
			"/login",
			"/signup",
			"/auth/refresh",
			"/payment-verify-webhook",
		}

		if !slices.Contains(allowed, r.URL.Path) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "No authorization token was provied", http.StatusUnauthorized)
				return
			}

			accessToken := strings.Split(authHeader, " ")[1]
			fmt.Println(accessToken)

			claims, err := authInterface.ValidateAccessToken(accessToken)
			if err != nil {
				http.Error(w, "Error validating access token", http.StatusUnauthorized)
				return
			}

			fmt.Println(claims)
		}

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"fmt"
	"net/http"
	"rayaw-api/internal/interfaces"
	"regexp"
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
			"/products",
		}

		// add dymanic path the allowed slice
		re, err := regexp.Compile(`(products/)\d+`)
		if err != nil {
			http.Error(w, fmt.Sprintf("error compiling regex: %v", err), http.StatusInternalServerError)
		}

		if re.MatchString(r.URL.Path) {
			allowed = append(allowed, r.URL.Path)
		}

		if !slices.Contains(allowed, r.URL.Path) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "No authorization token was provied", http.StatusUnauthorized)
				return
			}

			accessToken := strings.Split(authHeader, " ")[1]

			_, err := authInterface.ValidateAccessToken(accessToken)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error validating access token: %v", err), http.StatusUnauthorized)
				return
			}

		}

		next.ServeHTTP(w, r)
	})
}

package middlewares

import (
	"net/http"
	"simple-rest-api-golang/helper"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		_, err := helper.VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid Token!", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

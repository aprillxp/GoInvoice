package middleware

import (
	"api/utils"
	"net/http"
	"strings"

	"github.com/gorilla/context"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid or Expired token", http.StatusUnauthorized)
			return
		}

		context.Set(r, "user_id", claims["user_id"])
		next.ServeHTTP(w, r)
	})
}

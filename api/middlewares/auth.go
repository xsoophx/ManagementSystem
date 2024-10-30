package middlewares

import (
	"ManagementSystem/util"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := validateTokenAndGetUserID(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := util.SetUserIDInContext(r.Context(), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateTokenAndGetUserID(token string) (string, error) {
	return "12345", nil
}

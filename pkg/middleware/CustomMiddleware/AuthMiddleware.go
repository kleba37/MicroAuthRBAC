package CustomMiddleware

import (
	"net/http"
	"os"
	"strings"
)

type AuthMiddleware struct{}

func (m AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if !isNotEmptyToken(token) || !isBearer(token) || !isValidToken(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isNotEmptyToken(token string) bool {
	return len(token) > 0
}

func isValidToken(token string) bool {
	return os.Getenv("TOKEN") == token[7:]
}

func isBearer(token string) bool {
	return strings.HasPrefix(token, "Bearer")
}

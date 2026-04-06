package restapi

import (
	"log/slog"
	"net/http"
)

func loggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				
			next.ServeHTTP(w, r)
		})
	}
}

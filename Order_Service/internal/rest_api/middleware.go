package restapi

import (
	"log/slog"
	"net/http"
	"time"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *statusResponseWriter) Write(body []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}
	return w.ResponseWriter.Write(body)
}

func loggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			
			startedAt := time.Now()
			statusWriter := &statusResponseWriter{ResponseWriter: w}

			logger.Info("получен http-запрос",
				"method", r.Method,
				"path", r.URL.Path,
				"query", r.URL.RawQuery,
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)

			next.ServeHTTP(statusWriter, r)

			if statusWriter.statusCode == 0 {
				statusWriter.statusCode = http.StatusOK
			}

			logger.Info("обработан http-запрос",
				"method", r.Method,
				"path", r.URL.Path,
				"status_code", statusWriter.statusCode,
				"duration_ms", time.Since(startedAt).Milliseconds(),
			)
		})
	}
}

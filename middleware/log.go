package middleware

import (
	"log/slog"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			slog.Info("request received", "method", r.Method, "path", r.URL.Path)
			// Pass the request to the next middleware/handler in the chain
			next.ServeHTTP(w, r)
		},
	)
}

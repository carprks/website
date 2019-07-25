package backend

import (
	"net/http"
	"os"
)

func presetHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if os.Getenv("DEVELOPMENT") != "true" {
			w.Header().Set("Strict-Transport-Security", "max-age=1000; includeSubdomains; preload")
			w.Header().Set("Content-Security-Policy", "upgrade-insecure-requests")
			w.Header().Set("Feature-Policy", "vibrate 'none")
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
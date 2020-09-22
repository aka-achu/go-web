package middleware

import (
	"context"
	"github.com/aka-achu/go-web/logging"
	"github.com/google/uuid"
	"net/http"
)

// NoAuthLogging is a middleware without authentication
func NoAuthLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generating a unique id which will be used to trace the request
		traceID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "trace_id", traceID)

		logging.RequestLogger.Infof("TraceID-%s Addr-%s Protocol-%s Method-%s URI-%s",
			traceID, r.RemoteAddr, r.Proto, r.Method, r.RequestURI)
		next.ServeHTTP(
			w,
			r.WithContext(
				ctx,
			),
		)
	})
}

package middleware

import (
	"context"
	"github.com/aka-achu/go-web/logging"
	"github.com/google/uuid"
	"net/http"
)

func NoAuthLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.New().String()
		logging.RequestLogger.Infof("TraceID-%s Addr-%s Protocol-%s Method-%s URI-%s",
			traceID, r.RemoteAddr, r.Proto, r.Method, r.RequestURI)
		next.ServeHTTP(
			w,
			r.WithContext(
				context.WithValue(r.Context(), "trace_id", traceID),
			),
		)
	})
}

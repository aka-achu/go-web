package middleware

import (
	"context"
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/response"
	"github.com/aka-achu/go-web/utility"
	"github.com/google/uuid"
	"net/http"
)

// NoAuthLogging is a middleware which will omit the request authentication
// and generate a request_id for tracing the request and logs the request details
func NoAuthLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generating a unique id which will be used to trace the request
		traceID := uuid.New().String()
		// Storing the generated traceID in the context store against the key "trace_id"
		ctx := context.WithValue(r.Context(), "trace_id", traceID)
		logging.RequestLogger.Infof("TraceID-%s Addr-%s Protocol-%s Method-%s URI-%s",
			traceID, r.RemoteAddr, r.Proto, r.Method, r.RequestURI)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AuthLogging is a middleware which will validate the request based on the
// access_token token present in the request header and generate a request_id
// for tracing and logs the request details
func AuthLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generating a unique id which will be used to trace the request
		traceID := uuid.New().String()
		// Storing the generated traceID in the context store against the key "trace_id"
		ctx := context.WithValue(r.Context(), "trace_id", traceID)
		logging.RequestLogger.Infof("TraceID-%s Addr-%s Protocol-%s Method-%s URI-%s",
			traceID, r.RemoteAddr, r.Proto, r.Method, r.RequestURI)

		// Validating the Authorization token present in the request header
		if err := utility.VerifyToken(r.Header.Get("Authorization")); err != nil {
			logging.RequestLogger.Warnf("Invalid access_token. Error-%b TraceID-%s", err, traceID)
			response.UnAuthorized(w,"000",err.Error())
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
package middleware

import (
	"github.com/aka-achu/go-web/logging"
	"net/http"
)

func NoAuthLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.RequestLogger.Info(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.RequestURI)
		next.ServeHTTP(w, r)

	})
}

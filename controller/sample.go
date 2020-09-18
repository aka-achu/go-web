package controller

import (
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/response"
	"github.com/aka-achu/go-web/service"
	"net/http"
)

type Sample struct{}

var sampleService service.Sample

func (*Sample) HelloWorld(w http.ResponseWriter, r *http.Request) {
	requestTraceID := r.Context().Value("trace_id").(string)
	if message, err := sampleService.GetHostName(); err != nil {
		logging.AppLogger.Errorf("Failed to fetch the hostname of the system. TraceID-%s Error-%s",
			requestTraceID, err.Error())
		response.InternalServerError(w, "101", "Failed to get the hostname")
	} else {
		logging.AppLogger.Infof("Successfully fetched the hostname of the system. TraceID-%s", requestTraceID)
		response.Success(w, "100", "Successful sample request", message)
	}
}

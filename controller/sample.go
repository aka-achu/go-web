package controller

import (
	"github.com/aka-achu/go-web/response"
	"github.com/aka-achu/go-web/service"
	"net/http"
)

type Sample struct{}

var sampleService service.Sample

func (*Sample) HelloWorld(w http.ResponseWriter, r *http.Request) {
	if message, err := sampleService.GetHostName(); err != nil {
		response.InternalServerError(w, "101", "Failed to get the hostname")
	} else {
		response.Success(w, "100", "Successful sample request", message)
	}
}

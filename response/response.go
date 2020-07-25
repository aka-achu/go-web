package response

import (
	"encoding/json"
	"github.com/aka-achu/go-web/models"
	"net/http"
)

func getResponseBody(code string, message string, data ...interface{}) *models.Response {
	if len(data) == 0 {
		return &models.Response{
			Code:    code,
			Message: message,
		}
	} else {
		return &models.Response{
			Code:    code,
			Message: message,
			Data:    data[0],
		}
	}
}

func BadRequest(w http.ResponseWriter, code string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(getResponseBody(code, message))
}

func InternalServerError(w http.ResponseWriter, code string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(getResponseBody(code, message))
}

func Success(w http.ResponseWriter, code string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(getResponseBody(code, message, data))
}

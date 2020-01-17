package responses

import (
	"encoding/json"
	"net/http"
)

// Response struct
//TODO: make interface instead of map
type Response struct {
	Status  string                 `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"msg"`
}

// SendResponse send http response
func (entity Response) SendResponse(rw http.ResponseWriter, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")

	switch statusCode {
	case http.StatusOK:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
		entity.Message = "success"
	case http.StatusCreated:
		rw.WriteHeader(http.StatusCreated)
		entity.Status = http.StatusText(http.StatusCreated)
	case http.StatusAccepted:
		rw.WriteHeader(http.StatusAccepted)
		entity.Status = http.StatusText(http.StatusAccepted)
	case http.StatusServiceUnavailable:
		rw.WriteHeader(http.StatusServiceUnavailable)
		entity.Status = http.StatusText(http.StatusServiceUnavailable)
		entity.Message = "success"
	default:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
	}
	// send response
	json.NewEncoder(rw).Encode(entity)
	return
}

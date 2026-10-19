package utils

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Result  bool   `json:"result"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func SendApiResponse(w http.ResponseWriter, statusCode int, result bool, data any) {
	response := ApiResponse{
		Result: result,
	}

	if result != true {
		response.Message = data.(string)
	} else {
		response.Data = data
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"result": false, "message": "Error generating JSON response"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

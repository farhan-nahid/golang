package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Error string `json:"error"`
}

const (
	StatusOK = "OK"
	StatusError = "ERROR"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}


func WriteError(err error) ErrorResponse{
	return ErrorResponse{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidatorError(errors validator.ValidationErrors) ErrorResponse{	
	var errorMessage []string

	for _, err := range errors {
		switch err.ActualTag() {
		case "required":
			errorMessage = append(errorMessage, err.Field() + " is required")
		case "email":
			errorMessage = append(errorMessage, err.Field() + " is not a valid email")
		default:
			errorMessage = append(errorMessage, err.Field() + " is not valid")
		}
	
	}

	return ErrorResponse{
		Status: StatusError,
		Error:  strings.Join(errorMessage, ", "),
	}
}
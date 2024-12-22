package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/farhan-nahid/golang/students-api/internal/storage"
	"github.com/farhan-nahid/golang/students-api/internal/types"
	"github.com/farhan-nahid/golang/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)


func New (storage storage.Storage) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating new student")

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.WriteError(errors.New("request body is required")))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.WriteError(err))
			return
		}

		// validate request body
		if err := validator.New().Struct(student); err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.ValidatorError(err.(validator.ValidationErrors)))
			return
		}

		lastId, err:= storage.CreateStudent(student.FirstName, student.LastName, student.Age, student.Email)

		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.WriteError(err))
			return
		}



		response.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Student created successfully", "id": fmt.Sprintf("%d", lastId)})
		
	}
}
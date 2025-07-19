package main

import (
	"encoding/json"
	"errors"

	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var ErrMethodNotAllowed = errors.New("method is not allowed")
var ErrInvalideAge = errors.New("umur dibawah 18 tahun")
var ErrInvalidData = errors.New("email atau umur tidak boleh kosong")

type DataInput struct {
	email string
	age   string
}

type SuccessResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func validateData(data DataInput) error {
	email := strings.TrimSpace(data.email)
	if len(email) < 1 {
		return ErrInvalidData
	}

	ageString := strings.TrimSpace(data.age)
	age, err := strconv.Atoi(ageString)
	if err != nil || age < 18 {
		return ErrInvalidData
	}

	return nil

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		json.NewEncoder(w).Encode(ErrorResponse{Error: ErrMethodNotAllowed.Error()})
		return
	}
	rawEmail := r.URL.Query().Get("email")
	rawAge := r.URL.Query().Get("age")

	var data DataInput
	data.email = rawEmail
	data.age = rawAge

	errValidate := validateData(data)
	if errValidate != nil {

		errJson := json.NewEncoder(w).Encode(ErrorResponse{Error: errValidate.Error()})
		if errJson != nil {
			http.Error(w, errJson.Error(), http.StatusBadRequest)
			return
		}
		return
	}
	json.NewEncoder(w).Encode(SuccessResponse{Status: "ok"})

}

func main() {
	http.HandleFunc("/validate", handler)
	logrus.Info("Application started at port 8000")
	http.ListenAndServe(":8000", nil)
}

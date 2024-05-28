package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/djangbahevans/go-template/models"
)

func RespondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DecodeJSON(data io.ReadCloser, v interface{}) error {
	return json.NewDecoder(data).Decode(v)
}

func GetIDFromURL(r *http.Request) string {
	return r.PathValue("id")
}

func ValidateUser(user models.CreateUserRequest) error {
	if user.FirstName == "" {
		return ErrFirstNameRequired
	}

	if user.LastName == "" {
		return ErrLastNameRequired
	}

	if user.Email == "" {
		return ErrEmailRequired
	}

	if user.Password == "" {
		return ErrPasswordRequired
	}

	return nil
}

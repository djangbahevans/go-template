package handlers

import (
	"log"
	"net/http"

	"github.com/djangbahevans/go-template/models"
	"github.com/djangbahevans/go-template/services"
	"github.com/djangbahevans/go-template/utils"
)

type UserRoutes struct {
	userService services.UserService
}

func RegisterUserRoutes(r *http.ServeMux, userService services.UserService) {
	userHandler := &UserRoutes{userService}

	r.HandleFunc("GET /users", userHandler.GetUsers)
	r.HandleFunc("GET /users/{id}", userHandler.GetUser)
	r.HandleFunc("POST /users", userHandler.CreateUser)
	r.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	r.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)
}

func (u *UserRoutes) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.GetUsers()
	if err != nil {
		log.Printf("error getting users: %v", err)
		http.Error(w, "error getting users", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, users)
}

func (u *UserRoutes) GetUser(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIDFromURL(r)
	user, err := u.userService.GetUser(id)
	if err != nil {
		log.Printf("error getting user: %v", err)
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, user)
}

func (u *UserRoutes) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.CreateUserRequest
	err := utils.DecodeJSON(r.Body, &user)
	if err != nil {
		log.Printf("error decoding request body: %v", err)
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	newUser, err := u.userService.CreateUser(user)
	if err != nil {
		if err == utils.ErrEmailExists {
			http.Error(w, "email already exists", http.StatusBadRequest)
			return
		}
		log.Printf("error creating user: %v", err)
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, newUser)
}

func (u *UserRoutes) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIDFromURL(r)
	var user models.UpdateUserRequest
	err := utils.DecodeJSON(r.Body, &user)
	if err != nil {
		log.Printf("error decoding request body: %v", err)
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := u.userService.UpdateUser(id, user)
	if err != nil {
		log.Printf("error updating user: %v", err)
		http.Error(w, "error updating user", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, updatedUser)
}

func (u *UserRoutes) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIDFromURL(r)
	err := u.userService.DeleteUser(id)
	if err != nil {
		log.Printf("error deleting user: %v", err)
		http.Error(w, "error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

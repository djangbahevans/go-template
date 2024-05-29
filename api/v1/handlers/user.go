package handlers

import (
	"log"
	"net/http"

	"github.com/djangbahevans/go-template/models"
	"github.com/djangbahevans/go-template/services"
	"github.com/djangbahevans/go-template/utils"
)

type UserRoutes struct {
	userService services.IUserService
}

func NewUserRoutes(userService services.IUserService) *UserRoutes {
	return &UserRoutes{userService: userService}
}

func (u *UserRoutes) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /users", u.GetUsers)
	r.HandleFunc("GET /users/{id}", u.GetUser)
	r.HandleFunc("POST /users", u.CreateUser)
	r.HandleFunc("PUT /users/{id}", u.UpdateUser)
	r.HandleFunc("DELETE /users/{id}", u.DeleteUser)
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
	switch err {
	case nil:
	case utils.ErrUserNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	default:
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
	switch err {
	case nil:
	case utils.ErrEmailExists, utils.ErrFirstNameRequired, utils.ErrLastNameRequired, utils.ErrEmailRequired, utils.ErrPasswordRequired:
		http.Error(w, err.Error(), http.StatusConflict)
		return
	default:
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
	switch err {
	case nil:
	case utils.ErrUserNotFound, utils.ErrEmailExists, utils.ErrFirstNameRequired, utils.ErrLastNameRequired, utils.ErrEmailRequired, utils.ErrPasswordRequired:
		http.Error(w, err.Error(), http.StatusConflict)
		return
	default:
		log.Printf("error updating user: %v", err)
		http.Error(w, "error updating user", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, updatedUser)
}

func (u *UserRoutes) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIDFromURL(r)
	err := u.userService.DeleteUser(id)
	switch err {
	case nil:
	case utils.ErrUserNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	default:
		log.Printf("error deleting user: %v", err)
		http.Error(w, "error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

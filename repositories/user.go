package repositories

import "github.com/djangbahevans/go-template/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
}

package services

import (
	"github.com/djangbahevans/go-template/models"
	"github.com/djangbahevans/go-template/repositories"
	"github.com/djangbahevans/go-template/utils"
)

type IUserService interface {
	GetUsers() ([]models.UserResponse, error)
	GetUser(id string) (*models.UserResponse, error)
	CreateUser(user models.CreateUserRequest) (*models.UserResponse, error)
	UpdateUser(id string, user models.UpdateUserRequest) (*models.UserResponse, error)
	DeleteUser(id string) error
}

type userService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetUsers() ([]models.UserResponse, error) {
	users, err := s.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	userResponses := []models.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, *mapUserToUserResponse(&user))
	}

	return userResponses, nil
}

func (s *userService) GetUser(id string) (*models.UserResponse, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return mapUserToUserResponse(user), nil
}

func (s *userService) CreateUser(user models.CreateUserRequest) (*models.UserResponse, error) {
	err := validateCreateUser(user)
	if err != nil {
		return nil, err
	}

	u, err := s.userRepository.GetUserByEmail(user.Email)
	if u != nil || err == nil {
		return nil, utils.ErrEmailExists
	}

	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	newUser.PasswordHash = hashedPassword

	err = s.userRepository.CreateUser(&newUser)
	if err != nil {
		return nil, err
	}

	return mapUserToUserResponse(&newUser), nil
}

func (s *userService) UpdateUser(id string, user models.UpdateUserRequest) (*models.UserResponse, error) {
	err := validateUpdateUser(user)
	if err != nil {
		return nil, err
	}

	updatedUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	err = s.userRepository.UpdateUser(id, &updatedUser)
	if err != nil {
		return nil, err
	}

	return mapUserToUserResponse(&updatedUser), nil
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepository.DeleteUser(id)
}

func validateCreateUser(user models.CreateUserRequest) error {
	if user.FirstName == "" {
		return utils.ErrFirstNameRequired
	}

	if user.LastName == "" {
		return utils.ErrLastNameRequired
	}

	if user.Email == "" {
		return utils.ErrEmailRequired
	}

	if user.Password == "" {
		return utils.ErrPasswordRequired
	}

	return nil
}

func validateUpdateUser(user models.UpdateUserRequest) error {
	if user.FirstName == "" {
		return utils.ErrFirstNameRequired
	}

	if user.LastName == "" {
		return utils.ErrLastNameRequired
	}

	if user.Email == "" {
		return utils.ErrEmailRequired
	}

	return nil
}

func mapUserToUserResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

package inmemory

import (
	"sync"
	"time"

	"github.com/djangbahevans/go-template/models"
	"github.com/djangbahevans/go-template/utils"
)

type userRepository struct {
	mu    sync.RWMutex
	users map[string]*models.User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		users: make(map[string]*models.User),
	}
}

func (r *userRepository) CreateUser(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = utils.GenerateID()
	r.users[user.ID] = user

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return nil
}

func (r *userRepository) GetUser(id string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, utils.ErrUserNotFound
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, utils.ErrUserNotFound
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := []models.User{}
	for _, user := range r.users {
		users = append(users, *user)
	}

	return users, nil
}

func (r *userRepository) UpdateUser(id string, user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return utils.ErrUserNotFound
	}

	oldUser := r.users[id]

	user.ID = id
	user.UpdatedAt = time.Now()
	user.CreatedAt = oldUser.CreatedAt
	r.users[id] = user


	return nil
}

func (r *userRepository) DeleteUser(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return utils.ErrUserNotFound
	}

	delete(r.users, id)

	return nil
}

func (r *userRepository) Seed(users []models.User) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range users {
		r.users[user.ID] = &user
	}
}

func (r *userRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users = make(map[string]*models.User)
}

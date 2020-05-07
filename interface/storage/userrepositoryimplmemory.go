package storage

import (
	"errors"
	"sync"

	"github.com/DubrulleKevin/dkdirsrv/model"
)

// UserRepositoryImplMemory structure
type UserRepositoryImplMemory struct {
	mutex sync.Mutex
	users map[*model.IDModel]*model.UserModel
}

// NewUserRepositoryImplMemory function
func NewUserRepositoryImplMemory() *UserRepositoryImplMemory {
	return &UserRepositoryImplMemory{}
}

// GetAllUsers method
func (ur *UserRepositoryImplMemory) GetAllUsers() ([]*model.UserModel, error) {
	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	users := make([]*model.UserModel, len(ur.users))

	i := 0
	for _, user := range ur.users {
		users[i] = user
		i++
	}

	return users, nil
}

// GetUserByID method
func (ur *UserRepositoryImplMemory) GetUserByID(id *model.IDModel) (*model.UserModel, error) {
	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	user, ok := ur.users[id]
	if !ok {
		return nil, errors.New("No user for this ID")
	}

	return user, nil
}

// AddUser method
func (ur *UserRepositoryImplMemory) AddUser(id *model.IDModel, user *model.UserModel) error {
	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	_, ok := ur.users[id]

	if ok {
		return errors.New("ID already exists")
	}

	ur.users[id] = user
	return nil
}

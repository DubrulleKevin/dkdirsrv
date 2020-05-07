package repository

import (
	"github.com/DubrulleKevin/dkdirsrv/model"
)

// UserRepository interface
type UserRepository interface {
	GetAllUsers() ([]*model.UserModel, error)
	GetUserByID(id *model.IDModel) (*model.UserModel, error)
	AddUser(id *model.IDModel, user *model.UserModel) error
}

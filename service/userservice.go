package service

import "github.com/DubrulleKevin/dkdirsrv/repository"

// UserService structure
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService function
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

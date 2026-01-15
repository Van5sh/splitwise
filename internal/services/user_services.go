package services

import (
	"context"

	domain "github.com/Van5sh/new-splitwise/domain/repositories"
)

type UserServices struct {
	repo domain.UserRepository
}

func NewUsersServices(repo domain.UserRepository) *UserServices {
	return &UserServices{repo: repo}
}

func (s *UserServices) GetUsers(ctx context.Context)

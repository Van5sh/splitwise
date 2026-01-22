package services

import (
	"context"

	"github.com/Van5sh/new-splitwise/domain/models"
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
)

type UserServices struct {
	repo domain.UserRepository
}

func NewUsersServices(repo domain.UserRepository) *UserServices {
	return &UserServices{repo: repo}
}

func (s *UserServices) GetUsers(ctx context.Context) ([]models.User, error) {
	dbUsers, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]models.User, len(dbUsers))
	for i, u := range dbUsers {
		users[i] = models.User{
			ID:         u.ID,
			Role:       u.Role,
			FirebaseID: u.FirebaseUid,
			CreatedAt:  u.CreatedAt,
			UpdatedAt:  u.UpdatedAt,
		}
	}
	return users, nil
}

func (s *UserServices) GetUserByID(ctx context.Context, id string) (models.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		ID:         user.ID.String(),
		Role:       user.Role,
		FirebaseID: user.FirebaseUid,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}

func (s *UserServices) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return s.repo.CreateUser(ctx, user)
}

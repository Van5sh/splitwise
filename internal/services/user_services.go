package services

import (
	"context"
	"errors"

	"github.com/Van5sh/new-splitwise/domain/models"
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
	"github.com/google/uuid"
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
		ID:         user.ID,
		Role:       user.Role,
		FirebaseID: user.FirebaseUid,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}

func (s *UserServices) CreateUser(ctx context.Context, userName string, FirebaseID string, role string, email string) (models.User, error) {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return models.User{}, err
	}
	repo := s.repo.WithTx(tx)
	dbUser, err := repo.GetUserByName(ctx, userName)
	if err != nil {
		return models.User{}, err
	}
	if dbUser.ID != uuid.Nil {
		tx.Rollback()
		return models.User{}, errors.New("user already exists")
	}
	newUser, err := repo.CreateUser(ctx, FirebaseID, role)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}
	if err := tx.Commit(); err != nil {
		return models.User{}, err
	}
	return models.User{
		ID:         newUser.ID,
		Role:       newUser.Role,
		FirebaseID: newUser.FirebaseUid,
		CreatedAt:  newUser.CreatedAt,
		UpdatedAt:  newUser.UpdatedAt,
	}, nil
}

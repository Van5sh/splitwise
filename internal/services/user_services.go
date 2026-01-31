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

func (s *UserServices) GetUserByFirebaseID(ctx context.Context, firebaseID string) (models.User, error) {
	user, err := s.repo.GetUserByFirebaseID(ctx, firebaseID)
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

func (r *UserServices) UpdateUserRole(ctx context.Context, id string, role string) (models.User, error) {
	_, err := r.repo.ValidateUserExists(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	res, err := r.repo.UpdateUserRole(ctx, id, role)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		ID:         res.ID,
		Role:       res.Role,
		FirebaseID: res.FirebaseUid,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
	}, nil
}

func (r *UserServices) DeleteUser(ctx context.Context, id string) error {
	_, err := r.repo.ValidateUserExists(ctx, id)
	if err != nil {
		return err
	}
	_, err = r.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

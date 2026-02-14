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
			FirebaseID: u.FirebaseUid,
			UserName:   u.UserName,
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
		FirebaseID: user.FirebaseUid,
		UserName:   user.UserName,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}

func (s *UserServices) CreateUser(ctx context.Context, userName string, FirebaseID string, email string) (models.User, error) {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return models.User{}, err
	}
	repo := s.repo.WithTx(tx)
	newUser, err := repo.CreateUser(ctx, FirebaseID, userName)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}
	_, err = repo.CreateUserDetails(ctx, newUser.ID.String(), email)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}
	if err := tx.Commit(); err != nil {
		return models.User{}, err
	}
	return models.User{
		ID:         newUser.ID,
		FirebaseID: newUser.FirebaseUid,
		UserName:   newUser.UserName,
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
		FirebaseID: user.FirebaseUid,
		UserName:   user.UserName,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
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

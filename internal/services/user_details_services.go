package services

import (
	"context"

	"github.com/Van5sh/new-splitwise/domain/models"
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
)

type UserDetailsService struct {
	repo domain.UserDetailsRepository
}

func NewUserDetailsService(repo domain.UserDetailsRepository) *UserDetailsService {
	return &UserDetailsService{repo: repo}
}

func (s *UserDetailsService) GetUserDetailsByUserID(ctx context.Context, id string) (models.UserDetails, error) {
	res, err := s.repo.GetUserDetailsByUserID(ctx, id)
	if err != nil {
		return models.UserDetails{}, err
	}
	return models.UserDetails{
		ID:        res.ID,
		User_ID:   res.UserID,
		User_Name: res.UserName,
		Email:     res.Email,
	}, nil
}

func (S *UserDetailsService) GetUserDetailsByEmail(ctx context.Context, email string) (models.UserDetails, error) {
	res, err := S.repo.GetUserDetailsByEmail(ctx, email, "")
	if err != nil {
		return models.UserDetails{}, err
	}
	return models.UserDetails{
		ID:        res.ID,
		User_ID:   res.UserID,
		User_Name: res.UserName,
		Email:     res.Email,
	}, nil
}

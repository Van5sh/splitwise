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

func (s *UserDetailsService) GetUserDetailsByEmail(ctx context.Context, email string) (models.UserDetails, error) {
	res, err := s.repo.GetUserDetailsByEmail(ctx, email)
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

func (s *UserDetailsService) CheckUserInGroupParams(ctx context.Context, userId, groupId string) (int32, error) {
	return s.repo.CheckUserInGroup(ctx, userId, groupId)
}

func (s *UserDetailsService) CheckUserIsGroupAdmin(ctx context.Context, userId, groupId string) (int32, error) {
	return s.repo.CheckUserIsGroupAdmin(ctx, userId, groupId)
}

func (s *UserDetailsService) UpdateUserDetails(ctx context.Context, id, name string, email string) (models.UserDetails, error) {
	_, err := s.repo.GetUserDetailsByUserID(ctx, id)
	if err != nil {
		return models.UserDetails{}, err
	}
	res, err := s.repo.UpdateUserDetails(ctx, id, name, email)
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

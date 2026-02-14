package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Van5sh/new-splitwise/domain/models"
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
	"github.com/Van5sh/new-splitwise/internal/helpers"
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
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserDetails{}, helpers.NotFound("User details not found", nil, err)
		}
		return models.UserDetails{}, helpers.InternalServerError("Error fetching user details", nil, err)
	}
	return models.UserDetails{
		ID:        res.ID,
		User_ID:   res.UserID,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (s *UserDetailsService) GetUserDetailsByEmail(ctx context.Context, email string) (models.UserDetails, error) {
	res, err := s.repo.GetUserDetailsByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserDetails{}, helpers.NotFound("User details not found", nil, err)
		}
		return models.UserDetails{}, helpers.InternalServerError("Error fetching user details", nil, err)
	}
	return models.UserDetails{
		ID:        res.ID,
		User_ID:   res.UserID,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (s *UserDetailsService) CheckUserInGroupParams(ctx context.Context, userId, groupId string) (int32, error) {
	return s.repo.CheckUserInGroup(ctx, userId, groupId)
}

func (s *UserDetailsService) CheckUserIsGroupAdmin(ctx context.Context, userId, groupId string) (int32, error) {
	return s.repo.CheckUserIsGroupAdmin(ctx, userId, groupId)
}

func (s *UserDetailsService) UpdateUserDetails(ctx context.Context, id string, email string) (models.UserDetails, error) {
	_, err := s.repo.GetUserDetailsByUserID(ctx, id)
	if err != nil {
		return models.UserDetails{}, err
	}
	res, err := s.repo.UpdateUserDetails(ctx, id, email)
	if err != nil {
		return models.UserDetails{}, err
	}
	return models.UserDetails{
		ID:      res.ID,
		User_ID: res.UserID,
		Email:   res.Email,
	}, nil
}

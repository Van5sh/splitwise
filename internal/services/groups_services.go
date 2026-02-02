package services

import (
	"context"

	"github.com/Van5sh/new-splitwise/domain/models"
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
)

type GroupServices struct {
	repo domain.GroupsRepository
}

func NewGroupServices(repo domain.GroupsRepository) *GroupServices {
	return &GroupServices{repo: repo}
}

func (s *GroupServices) GetGroups(ctx context.Context) ([]models.Group, error) {
	dbUsers, err := s.repo.GetGroups(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]models.Group, len(dbUsers))
	for i, u := range dbUsers {
		users[i] = models.Group{
			ID:          u.ID,
			GroupName:   u.GroupName,
			Description: u.Description.String,
			TotalAmount: u.TotalAmount,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}
	}
	return users, nil
}

// TODO: implement other group services
// func (s *GroupServices) GetGroup

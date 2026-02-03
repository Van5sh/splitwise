package services

import (
	"context"
	"errors"

	"github.com/Van5sh/new-splitwise/domain/models"
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
	"github.com/google/uuid"
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

func (s *GroupServices) GetGroupById(ctx context.Context, id string) (models.Group, error) {
	status, err := s.repo.ValidateGroupExists(ctx, id)
	if err != nil {
		return models.Group{}, err
	}
	if status == false {
		return models.Group{}, errors.New("group does not exist")
	}
	dbGroup, err := s.repo.GetGroupById(ctx, id)
	if err != nil {
		return models.Group{}, err
	}
	return models.Group{
		ID:          dbGroup.ID,
		GroupName:   dbGroup.GroupName,
		Description: dbGroup.Description.String,
		TotalAmount: dbGroup.TotalAmount,
		CreatedAt:   dbGroup.CreatedAt,
		UpdatedAt:   dbGroup.UpdatedAt,
	}, nil
}

func (s *GroupServices) GetGroupByName(ctx context.Context, name string) (models.Group, error) {
	dbGroup, err := s.repo.GetGroupByName(ctx, name)
	if err != nil {
		return models.Group{}, err
	}
	return models.Group{
		ID:          dbGroup.ID,
		GroupName:   dbGroup.GroupName,
		Description: dbGroup.Description.String,
		TotalAmount: dbGroup.TotalAmount,
		CreatedAt:   dbGroup.CreatedAt,
		UpdatedAt:   dbGroup.UpdatedAt,
	}, nil
}

func (s *GroupServices) CreateGroup(ctx context.Context, groupName string, description string) (models.Group, error) {
	exGroup, err := s.repo.GetGroupByName(ctx, groupName)
	if err == nil && exGroup.ID != uuid.Nil {
		return models.Group{}, errors.New("group with the same name already exists")
	}
	dbGroup, err := s.repo.CreateGroup(ctx, groupName, description)
	if err != nil {
		return models.Group{}, err
	}
	return models.Group{
		ID:          dbGroup.ID,
		GroupName:   dbGroup.GroupName,
		Description: dbGroup.Description.String,
		TotalAmount: dbGroup.TotalAmount,
		CreatedAt:   dbGroup.CreatedAt,
		UpdatedAt:   dbGroup.UpdatedAt,
	}, nil
}

func (s *GroupServices) UpdateGroup(ctx context.Context, id string, groupName string, description string) (models.Group, error) {
	exGroup, err := s.repo.ValidateGroupExists(ctx, id)
	if err != nil {
		return models.Group{}, err
	}
	if exGroup == false {
		return models.Group{}, errors.New("group does not exist")
	}
	dbGroup, err := s.repo.UpdateGroup(ctx, id, groupName, description)
	if err != nil {
		return models.Group{}, err
	}
	return models.Group{
		ID:          dbGroup.ID,
		GroupName:   dbGroup.GroupName,
		Description: dbGroup.Description.String,
		TotalAmount: dbGroup.TotalAmount,
		CreatedAt:   dbGroup.CreatedAt,
		UpdatedAt:   dbGroup.UpdatedAt,
	}, nil
}

func (s *GroupServices) DeleteGroup(ctx context.Context, id string) error {
	exGroup, err := s.repo.ValidateGroupExists(ctx, id)
	if err != nil {
		return err
	}
	if exGroup == false {
		return errors.New("Group does not exist")
	}
	err = s.repo.DeleteGroup(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *GroupServices) GetGroupsByUserId(ctx context.Context, userID string) ([]models.Group, error) {
	{
		dbGroups, err := s.repo.GetGroupsByUserId(ctx, userID)
		if err != nil {
			return nil, err
		}
		groups := make([]models.Group, len(dbGroups))
		for i, g := range dbGroups {
			groups[i] = models.Group{
				ID:          g.ID,
				GroupName:   g.GroupName,
				Description: g.Description.String,
				TotalAmount: g.TotalAmount,
				CreatedAt:   g.CreatedAt,
				UpdatedAt:   g.UpdatedAt,
			}
		}
		return groups, nil
	}
}

func (s *GroupServices) GetGroupMembers(ctx context.Context, groupId string) ([]models.User, error) {
	members, err := s.repo.GetGroupMembers(ctx, groupId)
	if err != nil {
		return nil, err
	}
	users := make([]models.User, len(members))
	for i, u := range members {
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

func (s *GroupServices) GetGroupAdmins(ctx context.Context, groupId string) ([]models.User, error) {
	admins, err := s.repo.GetGroupAdmins(ctx, groupId)
	if err != nil {
		return nil, err
	}
	users := make([]models.User, len(admins))
	for i, u := range admins {
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

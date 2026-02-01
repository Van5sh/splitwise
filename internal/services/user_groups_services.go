package services

import (
	"context"

	domain "github.com/Van5sh/new-splitwise/domain/repositories"
)

type UserGroupService struct {
	repo domain.UserGroupRepository
}

func NewUserGroupService(repo domain.UserGroupRepository) *UserGroupService {
	return &UserGroupService{repo: repo}
}

func (s *UserGroupService) AddUserToGroup(ctx context.Context, userId, groupId string) error {
	return s.repo.AddUserToGroup(ctx, userId, groupId)
}

func (s *UserGroupService) RemoveUserFromGroup(ctx context.Context, userId, groupId string) error {
	return s.repo.RemoveUserFromGroup(ctx, userId, groupId)
}

func (s *UserGroupService) Close() error {
	return nil
}

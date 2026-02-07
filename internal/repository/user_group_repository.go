package repository

import (
	"context"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
)

type UserGroupRepository struct {
	q *sqlc.Queries
}

func NewUserGroupRepository(q *sqlc.Queries) *UserGroupRepository {
	return &UserGroupRepository{q: q}
}

func (r *UserGroupRepository) AddUserToGroup(ctx context.Context, userId, groupId string) error {
	userIdUUID, _ := uuid.Parse(userId)
	groupIdUUID, _ := uuid.Parse(groupId)
	return r.q.AddUserToGroup(ctx, sqlc.AddUserToGroupParams{
		UserID:  userIdUUID,
		GroupID: groupIdUUID,
	})
}

func (r *UserGroupRepository) RemoveUserFromGroup(ctx context.Context, userId, groupId string) error {
	userIdUUID, _ := uuid.Parse(userId)
	groupIdUUID, _ := uuid.Parse(groupId)
	return r.q.RemoveUserFromGroup(ctx, sqlc.RemoveUserFromGroupParams{
		UserID:  userIdUUID,
		GroupID: groupIdUUID,
	})
}

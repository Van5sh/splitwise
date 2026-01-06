package domain

import (
	"context"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
)

type GroupsRepository interface {
	GetGroups(ctx context.Context) ([]sqlc.Group, error)
	GetGroupById(ctx context.Context, id string) (sqlc.Group, error)
	GetGroupByName(ctx context.Context, name string) (sqlc.Group, error)
	CreateGroup(ctx context.Context, groupName string, description string) (sqlc.Group, error)
	UpdateGroup(ctx context.Context, id string, groupName string, description string) (sqlc.Group, error)
	DeleteGroup(ctx context.Context, id string) error
	GetGroupsByUserId(ctx context.Context, userId string) ([]sqlc.Group, error)
	GetGroupMembers(ctx context.Context, groupId string) ([]sqlc.User, error)
	GetGroupAdmins(ctx context.Context, groupId string) ([]sqlc.User, error)
	ValidateGroupExists(ctx context.Context, id string) (bool, error)
}

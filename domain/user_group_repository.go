package domain

import "context"

type UserGroupRepository interface {
	AddUserToGroup(ctx context.Context, userId, groupId string) error
	RemoveUserFromGroup(ctx context.Context, userId, groupId string) error
}

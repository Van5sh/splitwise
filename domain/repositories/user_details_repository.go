package domain

import (
	"context"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
)

type UserDetailsRepository interface {
	NewUserDetailsRepository(q *sqlc.Queries) *UserDetailsRepository
	GetUserDetailsByUserID(ctx context.Context, id string) (sqlc.UserDetail, error)
	GetUserDetailsByEmail(ctx context.Context, email string) (sqlc.UserDetail, error)
	UpdateUserDetails(ctx context.Context, id, name string, email string) (sqlc.UserDetail, error)
	CheckUserInGroupParams(ctx context.Context, userId, groupId string) (int32, error)
	CheckUserIsGroupAdmin(ctx context.Context, userId, groupId string) (int32, error)
}

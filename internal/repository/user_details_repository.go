package repository

import (
	"context"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
)

// new
type UserDetailsRepository struct {
	q *sqlc.Queries
}

func NewUserDetailsRepository(q *sqlc.Queries) *UserDetailsRepository {
	return &UserDetailsRepository{q: q}
}

func (r *UserDetailsRepository) GetUserDetailsByUserID(ctx context.Context, id string) (sqlc.GetUserDetailsByUserIdRow, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return sqlc.GetUserDetailsByUserIdRow{}, err
	}
	return r.q.GetUserDetailsByUserId(ctx, userID)
}

func (r *UserDetailsRepository) GetUserDetailsByEmail(ctx context.Context, email string) (sqlc.UserDetail, error) {
	return r.q.GetUserDetailsByEmail(ctx, email)
}

func (r *UserDetailsRepository) CreateUserDetails(ctx context.Context, userID string, name string, email string) (sqlc.UserDetail, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return sqlc.UserDetail{}, err
	}
	return r.q.CreateUserDetails(ctx, sqlc.CreateUserDetailsParams{
		UserID:   uid,
		UserName: name,
		Email:    email,
	})
}

func (r *UserDetailsRepository) UpdateUserDetails(ctx context.Context, id, UserName string, email string) (sqlc.UserDetail, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return sqlc.UserDetail{}, err
	}
	return r.q.UpdateUserDetails(ctx, sqlc.UpdateUserDetailsParams{
		UserName: UserName,
		Email:    email,
	})
}

func (r *UserDetailsRepository) CheckUserInGroup(ctx context.Context, userId, groupId string) (int32, error) {
	userIdUUID, _ := uuid.Parse(userId)
	groupIdUUID, _ := uuid.Parse(groupId)
	return r.q.CheckUserInGroup(ctx, sqlc.CheckUserInGroupParams{
		UserID:  userIdUUID,
		GroupID: groupIdUUID,
	})
}

func (r *UserDetailsRepository) CheckUserIsGroupAdmin(ctx context.Context, userId, groupId string) (int32, error) {
	userIdUUID, _ := uuid.Parse(userId)
	groupIdUUID, _ := uuid.Parse(groupId)
	return r.q.CheckUserIsGroupAdmin(ctx, sqlc.CheckUserIsGroupAdminParams{
		UserID:  userIdUUID,
		GroupID: groupIdUUID,
	})
}

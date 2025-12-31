package repository

import (
	"context"
	"database/sql"
	"errors"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
)

type GroupsRepository struct {
	q *sqlc.Queries
}

func NewGroupsRepository(q *sqlc.Queries) *GroupsRepository {
	return &GroupsRepository{q: q}
}

func (r *GroupsRepository) GetGroups(ctx context.Context) ([]sqlc.Group, error) {
	return r.q.GetGroups(ctx)
}

func (r *GroupsRepository) GetGroupById(ctx context.Context, id string) (sqlc.Group, error) {
	groupID, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Group{}, err
	}
	return r.q.GetGroupById(ctx, groupID)
}

func (r *GroupsRepository) GetGroupByName(ctx context.Context, name string) (sqlc.Group, error) {
	return r.q.GetGroupByName(ctx, name)
}

func (r *GroupsRepository) CreateGroup(
	ctx context.Context,
	groupName string,
	description string,
) (sqlc.Group, error) {
	return r.q.CreateGroup(ctx, sqlc.CreateGroupParams{
		GroupName: groupName,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
	})
}

func (r *GroupsRepository) UpdateGroup(
	ctx context.Context,
	id string,
	groupName string,
	description string,
) (sqlc.Group, error) {
	groupID, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Group{}, err
	}

	return r.q.UpdateGroup(ctx, sqlc.UpdateGroupParams{
		ID:        groupID,
		GroupName: groupName,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
	})
}

func (r *GroupsRepository) DeleteGroup(ctx context.Context, id string) (sqlc.Group, error) {
	groupID, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Group{}, err
	}
	return r.q.DeleteGroupById(ctx, groupID)
}

func (r *GroupsRepository) GetGroupsByUserId(
	ctx context.Context,
	userId string,
) ([]sqlc.Group, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	return r.q.GetGroupsByUserId(ctx, uid)
}

func (r *GroupsRepository) GetGroupMembers(
	ctx context.Context,
	id string,
) ([]sqlc.User, error) {
	groupID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.q.GetGroupMembers(ctx, groupID)
}

func (r *GroupsRepository) GetGroupAdmins(ctx context.Context, id string) ([]sqlc.User, error) {
	groupID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.q.GetGroupAdmins(ctx, groupID)
}

func (r *GroupsRepository) ValidateGroupExists(ctx context.Context, id string) (bool, error) {
	groupID, err := uuid.Parse(id)
	if err != nil {
		return false, err
	}

	_, err = r.q.ValidateGroupExists(ctx, groupID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

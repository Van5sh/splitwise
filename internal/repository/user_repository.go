package repository

import (
	"context"
	"database/sql"

	domain "github.com/Van5sh/new-splitwise/domain/repositories"
	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
	q  *sqlc.Queries
}

func NewUserRepository(db *sql.DB, q *sqlc.Queries) *UserRepository {
	return &UserRepository{db: db, q: q}
}

func (r *UserRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}

func (r *UserRepository) WithTx(tx *sql.Tx) domain.UserRepository {
	return &UserRepository{
		db: r.db,
		q:  r.q.WithTx(tx),
	}
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]sqlc.User, error) {
	return r.q.GetUsers(ctx)
}

func (r *UserRepository) GetUserByName(ctx context.Context, name string) (sqlc.User, error) {
	row, err := r.q.GetUserByName(ctx, name)
	if err != nil {
		return sqlc.User{}, err
	}
	return sqlc.User{
		ID:          row.ID,
		FirebaseUid: row.FirebaseUid,
		Role:        row.Role,
	}, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (sqlc.User, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return sqlc.User{}, nil
	}
	return r.q.GetUserById(ctx, userId)
}

func (r *UserRepository) CreateUser(ctx context.Context, firebaseId string, role string) (sqlc.User, error) {
	return r.q.CreateUser(ctx, sqlc.CreateUserParams{
		FirebaseUid: firebaseId,
		Role:        role,
	})
}

func (r *UserRepository) GetUserByFirebaseID(ctx context.Context, firebaseId string) (sqlc.User, error) {
	return r.q.GetUserByFirebaseUid(ctx, firebaseId)
}

func (r *UserRepository) UpdateUserRole(ctx context.Context, id string, role string) (sqlc.User, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return sqlc.User{}, err
	}
	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   userId,
		Role: role,
	})
}

func (r *UserRepository) CreateUserDetails(ctx context.Context, userID string, name string, email string) (sqlc.UserDetail, error) {
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

func (r *UserRepository) DeleteUser(ctx context.Context, id string) (sqlc.User, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return sqlc.User{}, err
	}
	err = r.q.DeleteUserById(ctx, userId)
	if err != nil {
		return sqlc.User{}, err
	}
	return sqlc.User{}, nil
}

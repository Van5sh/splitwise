package domain

import (
	"context"
	"database/sql"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
)

type UserRepository interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
	WithTx(tx *sql.Tx) UserRepository

	GetUsers(ctx context.Context) ([]sqlc.User, error)
	GetUserByID(ctx context.Context, id string) (sqlc.User, error)
	CreateUser(ctx context.Context, firebaseId string, role string) (sqlc.User, error)
	GetUserByName(ctx context.Context, name string) (sqlc.User, error)
	GetUserByFirebaseID(ctx context.Context, firebaseId string) (sqlc.User, error)
	UpdateUserRole(ctx context.Context, id string, role string) (sqlc.User, error)
	DeleteUser(ctx context.Context, id string) (sqlc.User, error)
	ValidateUserExists(ctx context.Context, id string) (bool, error)
}

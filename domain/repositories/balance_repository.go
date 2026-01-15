package domain

import (
	"context"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
)

type BalanceRepository interface {
	GetUserBalanceInGroup(ctx context.Context, userId string, groupId string) (int32, error)
	GetGroupBalances(ctx context.Context, id string) ([]sqlc.GetGroupBalanceRow, error)
}

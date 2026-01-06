package domain

import (
	"context"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
)

type SettlementRepository interface {
	GetSettlementById(ctx context.Context, id string) (sqlc.Settlement, error)
	AddSettlement(ctx context.Context, groupId, fromUserId, toUserId, amount string) (sqlc.Settlement, error)
	DeleteSettlementById(ctx context.Context, id string) (sqlc.Settlement, error)
	GetSettlementsByGroupId(ctx context.Context, groupId string) ([]sqlc.Settlement, error)
	GetSettlementsByUserId(ctx context.Context, userId string) ([]sqlc.Settlement, error)
	ValidateUsersInSameGroup(ctx context.Context, fromUserId, toUserId, groupId string) (int32, error)
}

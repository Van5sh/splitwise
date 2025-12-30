package repository

import (
	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
)

type BalanceRepository struct {
	q *sqlc.Queries
}

func NewBalanceRepository(q *sqlc.Queries) *BalanceRepository {
	return &BalanceRepository{q: q}
}

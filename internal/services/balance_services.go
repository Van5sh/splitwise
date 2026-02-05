package services

import (
	"context"

	"github.com/Van5sh/new-splitwise/domain/models"
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
)

type BalanceService struct {
	repo domain.BalanceRepository
}

func NewBalanceService(repo domain.BalanceRepository) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetUserBalanceInGroup(ctx context.Context, userID string, groupID string) (int32, error) {
	return s.repo.GetUserBalanceInGroup(ctx, userID, groupID)
}

func (s *BalanceService) GetGroupBalances(ctx context.Context, groupID string) ([]models.GetBalanceInGroup, error) {
	rows, err := s.repo.GetGroupBalances(ctx, groupID)
	if err != nil {
		return nil, err
	}

	balances := make([]models.GetBalanceInGroup, len(rows))
	for i, row := range rows {
		balances[i] = models.GetBalanceInGroup{
			UserID:  row.UserID,
			Balance: float64(row.Balance),
		}
	}
	return balances, nil
}

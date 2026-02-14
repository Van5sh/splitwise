package services

import (
	"context"

	"github.com/Van5sh/new-splitwise/domain/models"
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
)

type SettlementService struct {
	repo domain.SettlementRepository
}

func NewSettlementService(repo domain.SettlementRepository) *SettlementService {
	return &SettlementService{repo: repo}
}

func (s *SettlementService) GetSettlementById(ctx context.Context, id string) (models.Settlement, error) {
	res, err := s.repo.GetSettlementById(ctx, id)
	if err != nil {
		return models.Settlement{}, err
	}

	return models.Settlement{
		ID:         res.ID,
		GroupID:    res.GroupID,
		FromUserID: res.FromUserID,
		ToUserID:   res.ToUserID,
		Amount:     res.Amount,
		CreatedAt:  res.CreatedAt,
	}, nil
}

func (s *SettlementService) AddSettlement(ctx context.Context, groupId, fromUserId, toUserId, amount string) (models.Settlement, error) {
	_, err := s.repo.ValidateUsersInSameGroup(ctx, fromUserId, toUserId, groupId)
	if err != nil {
		return models.Settlement{}, err
	}
	res, err := s.repo.AddSettlement(ctx, groupId, fromUserId, toUserId, amount)
	if err != nil {
		return models.Settlement{}, err
	}
	return models.Settlement{
		ID:         res.ID,
		GroupID:    res.GroupID,
		FromUserID: res.FromUserID,
		ToUserID:   res.ToUserID,
		Amount:     res.Amount,
		CreatedAt:  res.CreatedAt,
	}, nil
}

func (s *SettlementService) DeleteSettlementById(ctx context.Context, id string) (models.Settlement, error) {
	_, err := s.repo.GetSettlementById(ctx, id)
	if err != nil {
		return models.Settlement{}, err
	}
	res, err := s.repo.DeleteSettlementById(ctx, id)
	if err != nil {
		return models.Settlement{}, err
	}
	return models.Settlement{
		ID:         res.ID,
		GroupID:    res.GroupID,
		FromUserID: res.FromUserID,
		ToUserID:   res.ToUserID,
		Amount:     res.Amount,
		CreatedAt:  res.CreatedAt,
	}, nil
}

func (s *SettlementService) GetSettlementsByGroupId(ctx context.Context, groupId string) ([]models.Settlement, error) {
	res, err := s.repo.GetSettlementsByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}
	var settlements []models.Settlement
	for _, settlement := range res {
		settlements = append(settlements, models.Settlement{
			ID:         settlement.ID,
			GroupID:    settlement.GroupID,
			FromUserID: settlement.FromUserID,
			ToUserID:   settlement.ToUserID,
			Amount:     settlement.Amount,
			CreatedAt:  settlement.CreatedAt,
		})
	}
	return settlements, nil
}

func (s *SettlementService) GetSettlementByUserID(ctx context.Context, userId string) ([]models.Settlement, error) {
	res, err := s.repo.GetSettlementsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	var settlements []models.Settlement
	for _, settlement := range res {
		settlements = append(settlements, models.Settlement{
			ID:         settlement.ID,
			GroupID:    settlement.GroupID,
			FromUserID: settlement.FromUserID,
			ToUserID:   settlement.ToUserID,
			Amount:     settlement.Amount,
			CreatedAt:  settlement.CreatedAt,
		})
	}
	return settlements, nil
}

func (s *SettlementService) ValidateUsersInSameGroup(ctx context.Context, fromUserId, toUserId, groupId string) (int32, error) {
	return s.repo.ValidateUsersInSameGroup(ctx, fromUserId, toUserId, groupId)
}

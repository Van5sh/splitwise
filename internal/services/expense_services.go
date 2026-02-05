package services

import (
	"context"

	"github.com/Van5sh/new-splitwise/domain/models"
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
	"github.com/gofiber/fiber/v2"
)

type ExpenseServices struct {
	repo domain.ExpenseRepository
}

func NewExpenseServices(repo domain.ExpenseRepository) *ExpenseServices {
	return &ExpenseServices{repo: repo}
}

func (s *ExpenseServices) GetExpenses(ctx context.Context) ([]models.Expense, error) {
	expenses, err := s.repo.GetExpenses(ctx)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	expense := make([]models.Expense, len(expenses))
	for i, u := range expenses {
		expense[i] = models.Expense{
			ID:          u.ID,
			GroupID:     u.GroupID,
			PaidBy:      u.PaidBy,
			Description: u.Description.String,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}
	}
	return expense, nil
}

func (s *ExpenseServices) GetExpenseByID(ctx context.Context, id string) (models.Expense, error) {
	expense, err := s.repo.GetExpenseByID(ctx, id)
	if err != nil {
		return models.Expense{}, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return models.Expense{
		ID:          expense.ID,
		GroupID:     expense.GroupID,
		PaidBy:      expense.PaidBy,
		Description: expense.Description.String,
		CreatedAt:   expense.CreatedAt,
		UpdatedAt:   expense.UpdatedAt,
	}, nil
}

func (s *ExpenseServices) GetExpenseByGroupID(ctx context.Context, id string) ([]models.Expense, error) {
	expenses, err := s.repo.GetExpenseByGroupID(ctx, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	expense := make([]models.Expense, len(expenses))
	for i, u := range expenses {
		expense[i] = models.Expense{
			ID:          u.ID,
			GroupID:     u.GroupID,
			PaidBy:      u.PaidBy,
			Description: u.Description.String,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}
	}
	return expense, nil
}

func (s *ExpenseServices) GetExpenseByUserId(ctx context.Context, id string) ([]models.Expense, error) {
	expenses, err := s.repo.GetExpenseByUserId(ctx, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	expense := make([]models.Expense, len(expenses))
	for i, u := range expenses {
		expense[i] = models.Expense{
			ID:          u.ID,
			GroupID:     u.GroupID,
			PaidBy:      u.PaidBy,
			Description: u.Description.String,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}
	}
	return expense, nil
}

func (s *ExpenseServices) CreateExpense(ctx context.Context, groupId, userId, description string, amount int) (models.Expense, error) {
	expense, err := s.repo.CreateExpense(ctx, groupId, userId, description, amount)
	if err != nil {
		return models.Expense{}, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return models.Expense{
		ID:          expense.ID,
		GroupID:     expense.GroupID,
		PaidBy:      expense.PaidBy,
		Description: expense.Description.String,
		CreatedAt:   expense.CreatedAt,
		UpdatedAt:   expense.UpdatedAt,
	}, nil
}

func (s *ExpenseServices) DeleteExpense(ctx context.Context, id string) error {
	err := s.repo.DeleteExpense(ctx, id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

func (s *ExpenseServices) ValidateExpenseExists(ctx context.Context, id string) (bool, error) {
	exists, err := s.repo.ValidateExpenseExists(ctx, id)
	if err != nil {
		return false, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return exists, nil
}

func (s *ExpenseServices) UpdateExpense(ctx context.Context, id, description string, amount int) (models.Expense, error) {
	expense, err := s.repo.UpdateExpense(ctx, id, description, amount)
	if err != nil {
		return models.Expense{}, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return models.Expense{
		ID:          expense.ID,
		GroupID:     expense.GroupID,
		PaidBy:      expense.PaidBy,
		Description: expense.Description.String,
		CreatedAt:   expense.CreatedAt,
		UpdatedAt:   expense.UpdatedAt,
	}, nil
}

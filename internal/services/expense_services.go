package services

import (
	domain "github.com/Van5sh/new-splitwise/domain/repositories"
)

type ExpenseServices struct {
	repo domain.ExpenseRepository
}

func NewExpenseServices(repo domain.ExpenseRepository) *ExpenseServices {
	return &ExpenseServices{repo: repo}
}

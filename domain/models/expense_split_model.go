package models

import (
	"time"

	"github.com/google/uuid"
)

type ExpenseSplit struct {
	ID        uuid.UUID `json:"id"`
	ExpenseID uuid.UUID `json:"expense_id"`
	UserID    uuid.UUID `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetBalanceInGroup struct {
	UserID  uuid.UUID `json:"user_id"`
	Balance float64   `json:"balance"`
}

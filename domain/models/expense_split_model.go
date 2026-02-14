package models

import (
	"time"

	"github.com/google/uuid"
)

type ExpenseSplit struct {
	ID        uuid.UUID `json:"id"`
	ExpenseID uuid.UUID `json:"expense_id"`
	UserID    uuid.UUID `json:"user_id"`
	PaidTo    uuid.UUID `json:"paid_to"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExpenseSplitInput struct {
	UserID string `json:"user_id"`
	PaidTo string `json:"paid_to"`
	Amount int    `json:"amount"`
}

type GetBalanceInGroup struct {
	UserID  uuid.UUID `json:"user_id"`
	Balance float64   `json:"balance"`
}

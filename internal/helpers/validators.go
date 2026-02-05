package helpers

import (
	"errors"

	"github.com/google/uuid"
)

type CreateExpenseRequest struct {
	GroupID string
	Amount  float64
	Note    string
}

func ValidateCreateExpense(req CreateExpenseRequest) error {
	if req.GroupID == "" {
		return errors.New("group_id is required")
	}

	if req.Amount <= 0 {
		return errors.New("amount must be > 0")
	}

	if len(req.Note) > 255 {
		return errors.New("note too long")
	}

	return nil
}

func ValidateId(id string) (uuid.UUID, error) {
	fid, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid id")
	}
	return fid, nil
}

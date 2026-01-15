package models

import "time"

type UserGroup struct {
	ID        string
	UserID    string
	GroupID   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

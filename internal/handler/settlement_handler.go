package handler

import (
	"github.com/Van5sh/new-splitwise/internal/services"
)

type SettlementHandler struct {
	services *services.SettlementService
}

func NewSettlementHandler(services *services.SettlementService) *SettlementHandler {
	return &SettlementHandler{services: services}
}

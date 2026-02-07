package handler

import (
	"github.com/Van5sh/new-splitwise/internal/helpers"
	"github.com/Van5sh/new-splitwise/internal/services"
	"github.com/gofiber/fiber/v2"
)

type BalanceHandler struct {
	services *services.BalanceService
}

func NewBalanceHandler(services *services.BalanceService) *BalanceHandler {
	return &BalanceHandler{services: services}
}

func (h *BalanceHandler) GetUserBalanceInGroup(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	groupID := c.Params("group_id")
	uid, err := helpers.ValidateId(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	balance, err := h.services.GetUserBalanceInGroup(c.Context(), uid.String(), groupID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"balance": balance})
}

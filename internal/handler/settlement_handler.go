package handler

import (
	"github.com/Van5sh/new-splitwise/internal/services"
	"github.com/gofiber/fiber/v2"
)

type SettlementHandler struct {
	services *services.SettlementService
}

func NewSettlementHandler(services *services.SettlementService) *SettlementHandler {
	return &SettlementHandler{services: services}
}

func (h *SettlementHandler) GetSettlementById(c *fiber.Ctx) error {
	id := c.Params("id")
	settlement, err := h.services.GetSettlementById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"settlement": settlement})
}

func (h *SettlementHandler) AddSettlement(c *fiber.Ctx) error {
	type request struct {
		GroupID    string `json:"group_id"`
		FromUserID string `json:"from_user_id"`
		ToUserID   string `json:"to_user_id"`
		Amount     string `json:"amount"`
	}
	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	settlement, err := h.services.AddSettlement(c.Context(), body.GroupID, body.FromUserID, body.ToUserID, body.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"settlement": settlement})
}

func (h *SettlementHandler) DeleteSettlementById(c *fiber.Ctx) error {
	id := c.Params("id")
	settlement, err := h.services.DeleteSettlementById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"settlement": settlement})
}

func (h *SettlementHandler) GetSettlementsByGroupId(c *fiber.Ctx) error {
	groupId := c.Params("group_id")
	settlements, err := h.services.GetSettlementsByGroupId(c.Context(), groupId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"settlements": settlements})
}

func (h *SettlementHandler) GetSettlementsByUserId(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	settlements, err := h.services.GetSettlementByUserID(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"settlements": settlements})
}

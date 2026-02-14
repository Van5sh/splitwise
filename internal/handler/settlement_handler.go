package handler

import (
	"encoding/json"

	"github.com/Van5sh/new-splitwise/internal/helpers"
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
	var payload map[string]interface{}
	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	groupID := getStringField(payload, "group_id", "groupId")
	fromUserID := getStringField(payload, "from_user_id", "fromUserId", "paid_by", "paidBy")
	toUserID := getStringField(payload, "to_user_id", "toUserId", "paid_to", "paidTo")
	amount, err := getIntField(payload, "amount", "Amount")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid amount"})
	}

	if groupID == "" || fromUserID == "" || toUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "group_id, from_user_id and to_user_id are required"})
	}
	if fromUserID == toUserID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "from_user_id and to_user_id must be different"})
	}
	if amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "amount must be greater than 0"})
	}
	if _, err := helpers.ValidateID(groupID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group id"})
	}
	if _, err := helpers.ValidateID(fromUserID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid from_user_id"})
	}
	if _, err := helpers.ValidateID(toUserID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid to_user_id"})
	}

	settlement, err := h.services.AddSettlement(c.Context(), groupID, fromUserID, toUserID, amount)
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

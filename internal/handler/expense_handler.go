package handler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Van5sh/new-splitwise/internal/helpers"
	"github.com/Van5sh/new-splitwise/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ExpenseHandler struct {
	services *services.ExpenseServices
}

func NewExpenseHandler(services *services.ExpenseServices) *ExpenseHandler {
	return &ExpenseHandler{services: services}
}

func (h *ExpenseHandler) GetExpense(c *fiber.Ctx) error {
	expenses, err := h.services.GetExpenses(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(expenses)
}

func (h *ExpenseHandler) GetExpenseByID(c *fiber.Ctx) error {
	id := c.Params("id")
	expense, err := h.services.GetExpenseByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	splits, err := h.services.GetExpenseSplitsByExpenseID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"expense": expense, "splits": splits})
}

func (h *ExpenseHandler) GetExpenseByGroupID(c *fiber.Ctx) error {
	groupId := c.Params("group_id")
	if _, err := helpers.ValidateID(groupId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group id"})
	}
	expenses, err := h.services.GetExpenseByGroupID(c.Context(), groupId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(expenses)
}

func (h *ExpenseHandler) CreateExpense(c *fiber.Ctx) error {
	groupId := c.Params("id")
	if _, err := helpers.ValidateID(groupId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group id"})
	}
	var payload map[string]interface{}
	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	paidBy := getStringField(payload, "paid_by", "paidBy")
	description := getStringField(payload, "description", "Description")
	amount, err := getIntField(payload, "amount", "Amount")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid amount"})
	}
	if paidBy == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "paid_by is required"})
	}
	if _, err := helpers.ValidateID(paidBy); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid paid_by user id"})
	}
	expense, err := h.services.CreateExpense(c.Context(), groupId, paidBy, description, amount, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(expense)
}

func (h *ExpenseHandler) CreateIndividualExpense(c *fiber.Ctx) error {
	var payload map[string]interface{}
	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	fromUserID := getStringField(payload, "from_user_id", "fromUserId", "paid_by", "paidBy")
	toUserID := getStringField(payload, "to_user_id", "toUserId")
	description := getStringField(payload, "description", "Description")
	amount, err := getIntField(payload, "amount", "Amount")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid amount"})
	}
	if fromUserID == "" || toUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "from_user_id and to_user_id are required"})
	}
	if fromUserID == toUserID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "from_user_id and to_user_id must be different"})
	}
	if _, err := helpers.ValidateID(fromUserID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid from_user_id"})
	}
	if _, err := helpers.ValidateID(toUserID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid to_user_id"})
	}

	expense, err := h.services.CreateIndividualExpense(c.Context(), fromUserID, toUserID, description, amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(expense)
}

func getStringField(payload map[string]interface{}, keys ...string) string {
	for _, key := range keys {
		if v, ok := payload[key]; ok {
			switch val := v.(type) {
			case string:
				return strings.TrimSpace(val)
			default:
				return strings.TrimSpace(fmt.Sprint(val))
			}
		}
	}
	return ""
}

func getIntField(payload map[string]interface{}, keys ...string) (int, error) {
	for _, key := range keys {
		if v, ok := payload[key]; ok {
			switch val := v.(type) {
			case float64:
				return int(val), nil
			case string:
				return strconv.Atoi(strings.TrimSpace(val))
			default:
				return 0, fmt.Errorf("invalid number")
			}
		}
	}
	return 0, fmt.Errorf("missing number")
}

func (h *ExpenseHandler) DeleteExpenseByID(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := helpers.ValidateID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid expense id"})
	}
	err = h.services.DeleteExpense(c.Context(), uid.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Expense deleted successfully"})
}

func (h *ExpenseHandler) UpdateExpense(c *fiber.Ctx) error {
	var payload map[string]interface{}
	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	id := c.Params("id")
	uid, err := helpers.ValidateID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid expense id"})
	}
	description := getStringField(payload, "description", "Description")
	amount, err := getIntField(payload, "amount", "Amount")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid amount"})
	}
	updatedExpense, err := h.services.UpdateExpense(c.Context(), uid.String(), description, amount, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Expense updated successfully", "expense": updatedExpense})
}

func (h *ExpenseHandler) ValidateExpenseExists(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := helpers.ValidateID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid expense id"})
	}
	exists, err := h.services.ValidateExpenseExists(c.Context(), uid.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"exists": exists})
}

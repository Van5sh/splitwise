package handler

import (
	"strconv"

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
	return c.Status(fiber.StatusOK).JSON(expense)
}

func (h *ExpenseHandler) GetExpenseByGroupID(c *fiber.Ctx) error {
	groupId := c.Params("group_id")
	expenses, err := h.services.GetExpenseByGroupID(c.Context(), groupId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(expenses)
}

func (h *ExpenseHandler) CreateExpense(c *fiber.Ctx) error {
	type request struct {
		GroupID     string `json:"group_id"`
		PaidBy      string `json:"paid_by"`
		Description string `json:"description"`
		Amount      string `json:"amount"`
	}
	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	amount, err := strconv.Atoi(body.Amount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid amount"})
	}
	expense, err := h.services.CreateExpense(c.Context(), body.GroupID, body.PaidBy, body.Description, amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(expense)
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
	type request struct {
		Description string `json:"description"`
		Amount      string `json:"amount"`
	}
	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	id := c.Params("id")
	uid, err := helpers.ValidateID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid expense id"})
	}
	amount, err := strconv.Atoi(body.Amount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid amount"})
	}
	expense, err := h.services.UpdateExpense(c.Context(), uid.String(), body.Description, amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Expense updated successfully", "expense": expense})
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

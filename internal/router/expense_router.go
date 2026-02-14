package router

import (
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type ExpenseRouter struct {
	handler *handler.ExpenseHandler
}

func NewExpenseRouter(handler *handler.ExpenseHandler) *ExpenseRouter {
	return &ExpenseRouter{handler: handler}
}

func (r *ExpenseRouter) ExpenseRouters(app *fiber.App) {
	expense := app.Group("/expense")
	expense.Get("/", r.handler.GetExpense)
	expense.Get("/:group_id", r.handler.GetExpenseByGroupID)
	expense.Get("/:id", r.handler.GetExpenseByID)
	expense.Delete("/:id", r.handler.DeleteExpenseByID)
	expense.Post("/:id", r.handler.CreateExpense)
	expense.Patch("/:id", r.handler.UpdateExpense)
}

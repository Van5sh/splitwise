package router

import (
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type BalanceRouter struct {
	handler *handler.BalanceHandler
}

func NewBalanceRouter(handler *handler.BalanceHandler) *BalanceRouter {
	return &BalanceRouter{handler: handler}
}

func (r *BalanceRouter) BalanceRouters(app *fiber.App) {
	balance := app.Group("/balance")
	balance.Get("/group/:group_id", r.handler.GetGroupBalances)
	balance.Get("/user/:user_id/group/:group_id", r.handler.GetUserBalanceInGroup)
}

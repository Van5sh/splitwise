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

}

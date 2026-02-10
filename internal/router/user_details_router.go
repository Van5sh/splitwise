package router

import (
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type UserDetailsRouter struct {
	handler *handler.BalanceHandler
}

func NewUserDetailsRouter(handler *handler.BalanceHandler) *BalanceRouter {
	return &BalanceRouter{handler: handler}
}

func (r *BalanceRouter) UserDetailsRoutes(app *fiber.App) {
}

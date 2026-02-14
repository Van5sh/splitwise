package router

import (
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type SettlementRouter struct {
	handler *handler.SettlementHandler
}

func NewSettlementRouter(handler *handler.SettlementHandler) *SettlementRouter {
	return &SettlementRouter{handler: handler}
}

func (r *SettlementRouter) SettlementRouters(app *fiber.App) {
	settlement := app.Group("/settlement")
	settlement.Get("/:id", r.handler.GetSettlementById)
	settlement.Get("/group/:group_id", r.handler.GetSettlementsByGroupId)
	settlement.Get("/user/:user_id", r.handler.GetSettlementsByUserId)
	settlement.Post("/add", r.handler.AddSettlement)
	settlement.Delete("/:id", r.handler.DeleteSettlementById)
}

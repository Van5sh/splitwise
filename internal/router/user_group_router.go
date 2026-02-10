package router

import (
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type UserGroupRouter struct {
	handler *handler.UserGroupHandler
}

func NewUserGroupRouter(handler *handler.UserGroupHandler) *UserGroupRouter {
	return &UserGroupRouter{handler: handler}
}

func (h *UserGroupRouter) UserGroupRoutes(app fiber.App) {

}

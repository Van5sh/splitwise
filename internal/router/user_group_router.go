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

func (r *UserGroupRouter) UserGroupRoutes(app *fiber.App) {
	ugroup := app.Group("/groups")
	ugroup.Post("/add", r.handler.AddUserToGroup)
	ugroup.Post("/remove", r.handler.RemoveUserFromGroup)
}

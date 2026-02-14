package router

import (
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type GroupRouter struct {
	handler *handler.GroupHandler
}

func NewGroupRouter(handler *handler.GroupHandler) *GroupRouter {
	return &GroupRouter{handler: handler}
}

func (r *GroupRouter) GroupRouters(app *fiber.App) {
	groups := app.Group("/groups")
	groups.Get("/", r.handler.GetGroups)
	groups.Get("/user/:user_id", r.handler.GetGroupsByUserId)
	groups.Get("/members/:group_id", r.handler.GetGroupMembers)
	groups.Get("/admins/:group_id", r.handler.GetGroupAdmins)
	groups.Get("/name/:name", r.handler.GetGroupByName)
	groups.Get("/:id", r.handler.GetGroupById)
	groups.Post("/new", r.handler.CreateGroup)
	groups.Patch("/:id", r.handler.UpdateGroup)
	groups.Delete("/:id", r.handler.DeleteGroup)
}

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
	groups.Get("", r.handler.GetGroups)
	groups.Get(":user_Id", r.handler.GetGroupsByUserId)
	groups.Get("/members/:group_id", r.handler.GetGroupMembers)
	groups.Get("/admins/:group_id", r.handler.GetGroupAdmins)
	groups.Get(":group_id", r.handler.GetGroupById)
	groups.Post("/new", r.handler.CreateGroup)
	groups.Patch("/change", r.handler.UpdateGroup)
	groups.Delete("/", r.handler.DeleteGroup)
}

package router

import (
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type UserDetailsRouter struct {
	handler *handler.UserDetailsHandler
}

func NewUserDetailsRouter(handler *handler.UserDetailsHandler) *UserDetailsRouter {
	return &UserDetailsRouter{handler: handler}
}

func (r *UserDetailsRouter) UserDetailsRoutes(app *fiber.App) {
	details := app.Group("/userdetails")

	details.Get(":email", r.handler.GetUserDetailsByEmail)
	details.Get("/:id", r.handler.GetUserDetails)
	details.Get("/check", r.handler.CheckUserInGroup)
	details.Patch("/new", r.handler.UpdateUserDetails)
}

package router

import (
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	handler *handler.UserHandler
}

func NewUserRouter(handler *handler.UserHandler) *UserRouter {
	return &UserRouter{handler: handler}
}

func (r *UserRouter) RegisterUserRoutes(app *fiber.App) {
	user := app.Group("/users")

	user.Get("/", r.handler.GetUsers)
	user.Get("/user/:id", r.handler.GetUserId)
	user.Get("/user/firebase/:firebase_id", r.handler.GetUserByFirebaseID)
	user.Post("/user", r.handler.CreateUser)
	user.Delete("/user/:id", r.handler.DeleteUser)
}

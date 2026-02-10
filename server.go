package main

import (
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/Van5sh/new-splitwise/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func startServer() {
	app := fiber.New()
	app.Use(logger.New())
	// TODO: Add CORS middleware
	// TODO: Add authentication middleware
	// app.User(middleware.Cors)

	// Initialize routers
	userRouter := router.NewUserRouter(&handler.UserHandler{})
	userDetailsRouter := router.NewUserDetailsRouter(&handler.UserDetailsHandler{})
	userGroupRouter := router.NewUserGroupRouter(&handler.UserGroupHandler{})
	settlementRouter := router.NewSettlementRouter(&handler.SettlementHandler{})
	expenseRouter := router.NewExpenseRouter(&handler.ExpenseHandler{})
	groupRouter := router.NewGroupRouter(&handler.GroupHandler{})
	balanceRouter := router.NewBalanceRouter(&handler.BalanceHandler{})

	// Register routes
	userRouter.RegisterUserRoutes(app)
	userDetailsRouter.UserDetailsRoutes(app)
	expenseRouter.ExpenseRouters(app)
	userGroupRouter.UserGroupRoutes(app)
	groupRouter.GroupRouters(app)
	balanceRouter.BalanceRouters(app)
	settlementRouter.SettlementRouters(app)

	// Start the server
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

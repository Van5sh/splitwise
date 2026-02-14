package main

import (
	"database/sql"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/Van5sh/new-splitwise/internal/handler"
	"github.com/Van5sh/new-splitwise/internal/middleware"
	"github.com/Van5sh/new-splitwise/internal/repository"
	"github.com/Van5sh/new-splitwise/internal/router"
	"github.com/Van5sh/new-splitwise/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func startServer(dbConn *sql.DB) {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(middleware.RequestTimingMiddleware())
	app.Use(middleware.ErrorResponseMiddleware())
	// TODO: Add CORS middleware
	// TODO: Add authentication middleware

	queries := sqlc.New(dbConn)
	userRepo := repository.NewUserRepository(dbConn, queries)
	userDetailsRepo := repository.NewUserDetailsRepository(queries)
	userGroupRepo := repository.NewUserGroupRepository(queries)
	settlementRepo := repository.NewSettlementRepository(queries)
	expenseRepo := repository.NewExpenseRepository(dbConn, queries)
	groupRepo := repository.NewGroupsRepository(queries)
	balanceRepo := repository.NewBalanceRepository(queries)

	userService := services.NewUsersServices(userRepo)
	userDetailsService := services.NewUserDetailsService(userDetailsRepo)
	userGroupService := services.NewUserGroupService(userGroupRepo)
	settlementService := services.NewSettlementService(settlementRepo)
	expenseService := services.NewExpenseServices(expenseRepo)
	groupService := services.NewGroupServices(groupRepo)
	balanceService := services.NewBalanceService(balanceRepo)

	userHandler := handler.NewUserHandler(userService)
	userDetailsHandler := handler.NewUserDetailsHandler(userDetailsService)
	userGroupHandler := handler.NewUserGroupHandler(userGroupService)
	settlementHandler := handler.NewSettlementHandler(settlementService)
	expenseHandler := handler.NewExpenseHandler(expenseService)
	groupHandler := handler.NewGroupHandler(groupService)
	balanceHandler := handler.NewBalanceHandler(balanceService)

	// Initialize routers
	userRouter := router.NewUserRouter(userHandler)
	userDetailsRouter := router.NewUserDetailsRouter(userDetailsHandler)
	userGroupRouter := router.NewUserGroupRouter(userGroupHandler)
	settlementRouter := router.NewSettlementRouter(settlementHandler)
	expenseRouter := router.NewExpenseRouter(expenseHandler)
	groupRouter := router.NewGroupRouter(groupHandler)
	balanceRouter := router.NewBalanceRouter(balanceHandler)

	// Register routes
	userRouter.RegisterUserRoutes(app)
	userDetailsRouter.UserDetailsRoutes(app)
	expenseRouter.ExpenseRouters(app)
	userGroupRouter.UserGroupRoutes(app)
	groupRouter.GroupRouters(app)
	balanceRouter.BalanceRouters(app)
	settlementRouter.SettlementRouters(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

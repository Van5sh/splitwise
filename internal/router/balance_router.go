package router

import "github.com/Van5sh/new-splitwise/internal/handler"

type BalanceRouter struct {
	handler *handler.BalanceHandler
}

func NewBalanceRouter(handler *handler.BalanceHandler) *BalanceRouter {
	return {handler:handler}
}
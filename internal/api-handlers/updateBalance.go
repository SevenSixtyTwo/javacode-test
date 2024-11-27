package apihandlers

import (
	structs "javacode-test/api/structs"
	ctxvalue "javacode-test/internal/ctx-value"
	dbhandlers "javacode-test/internal/db-handlers"
	"javacode-test/util/workerpool"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func UpdateBalance(c echo.Context) error {
	cc := c.(*structs.CustomContext)
	log := ctxvalue.GetLog(cc.Ctx)
	wp := ctxvalue.GetWP(cc.Ctx)

	account := &structs.Wallet{}
	if err := c.Bind(account); err != nil {
		log.Error("binding", "error", err)
		return c.JSON(http.StatusBadRequest, account)
	}

	var t *workerpool.Task

	if account.Operation == "WITHDRAW" {
		t = workerpool.NewTask(func() error {
			return dbhandlers.Withdraw(cc.Ctx, account.ID, account.Balance)
		}, nil)
	} else if account.Operation == "DEPOSIT" {
		t = workerpool.NewTask(func() error {
			return dbhandlers.Deposit(cc.Ctx, account.ID, account.Balance)
		}, nil)
	} else {
		return c.JSON(http.StatusBadRequest, "invalid operation")
	}

	wp.AddTask(t)

	if err := <-t.Err; err != nil {
		log.Error("update balance", "error", err)
		if strings.Contains(err.Error(), "insufficient funds") {
			return c.JSON(http.StatusBadRequest, "insufficient funds or wrong uuid")
		}
		if strings.Contains(err.Error(), "wrong uuid") {
			return c.JSON(http.StatusBadRequest, "wrong uuid")
		}
		return c.JSON(http.StatusInternalServerError, account)
	}

	return c.JSON(http.StatusOK, account)
}

package main

import (
	"context"
	"errors"
	"fmt"
	ctxvalue "javacode-test/internal/ctx-value"
	"javacode-test/internal/db"
	dbhandlers "javacode-test/internal/db-handlers"
	"javacode-test/internal/logger"
	"javacode-test/util/workerpool"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Wallet struct {
	ID        uuid.UUID `json:"valletid"`
	Operation string    `json:"operationType"`
	Balance   float64   `json:"amount"`
}

type CustomContext struct {
	echo.Context
	ctx context.Context
}

func updateBalance(c echo.Context) error {
	cc := c.(*CustomContext)
	log := ctxvalue.GetLog(cc.ctx)
	db := ctxvalue.GetDbPostgres(cc.ctx)
	wp := ctxvalue.GetWP(cc.ctx)

	account := &Wallet{}
	if err := c.Bind(account); err != nil {
		log.Error("binding", "error", err)
		return c.JSON(http.StatusBadRequest, account)
	}

	t := workerpool.NewTask(func() error {
		if account.Operation == "WITHDRAW" {
			return dbhandlers.Withdraw(cc.ctx, db, log, account.ID, account.Balance)
		} else if account.Operation == "DEPOSIT" {
			return dbhandlers.Deposit(cc.ctx, db, log, account.ID, account.Balance)
		} else {
			return fmt.Errorf("invalid operation")
		}
	}, nil)

	wp.AddTask(t)

	if err := <-t.Err; err != nil {
		log.Error("update balance", "error", err)
		if strings.Contains(err.Error(), "insufficient funds") {
			return c.JSON(http.StatusBadRequest, "insufficient funds")
		} else if strings.Contains(err.Error(), "invalid operation") {
			return c.JSON(http.StatusBadRequest, "invalid operation")
		}
		return c.JSON(http.StatusInternalServerError, account)
	}

	return c.JSON(http.StatusOK, account)
}

func getBalance(c echo.Context) error {
	cc := c.(*CustomContext)
	db := ctxvalue.GetDbPostgres(cc.ctx)
	log := ctxvalue.GetLog(cc.ctx)
	id := uuid.MustParse(c.Param("uuid"))

	balance, err := dbhandlers.GetBalance(cc.ctx, db, id)
	if err != nil {
		log.Error("get balance from db", "error", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	responseWallet := Wallet{id, "balance", balance}

	return c.JSON(http.StatusOK, responseWallet)
}

func main() {
	ctx := context.Background()
	log := logger.SetupLogger()

	log.Info("starting bank api")

	dbp, err := db.GetPostgresDb(ctx)
	if err != nil {
		log.Error("db postgres", "error", err)
		panic(fmt.Errorf("get db postgres %v", err))
	}

	wp := workerpool.NewPool(40, 400)

	wp.RunBackground()
	defer wp.Stop()

	ctx = context.WithValue(ctx, ctxvalue.ValueDbPostgres, dbp)
	ctx = context.WithValue(ctx, ctxvalue.ValueLog, log)
	ctx = context.WithValue(ctx, ctxvalue.ValueWP, wp)

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, ctx}
			return next(cc)
		}
	})
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "request timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			log.Error("request timeout", "error", c.Path())
		},
		Timeout: time.Second * 20,
	}))

	e.POST("/api/v1/wallet", updateBalance)

	e.GET("api/v1/wallets/:uuid", getBalance)

	if err := e.Start(":3030"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("failed to start server", "error", err)
		panic(fmt.Errorf("failed to start server %v", err))
	}
}

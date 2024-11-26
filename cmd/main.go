package main

import (
	"context"
	"errors"
	"fmt"
	"javacode-test/api/structs"
	apihanlders "javacode-test/internal/api-handlers"
	ctxvalue "javacode-test/internal/ctx-value"
	"javacode-test/internal/db"
	"javacode-test/internal/logger"
	"javacode-test/util/workerpool"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
			cc := &structs.CustomContext{c, ctx}
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

	e.POST("/api/v1/wallet", apihanlders.UpdateBalance)

	e.GET("/api/v1/wallets/:uuid", apihanlders.GetBalance)

	if err := e.Start(":3030"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("failed to start server", "error", err)
		panic(fmt.Errorf("failed to start server %v", err))
	}
}

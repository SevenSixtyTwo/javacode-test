package apihandlers

import (
	structs "javacode-test/api/structs"
	ctxvalue "javacode-test/internal/ctx-value"
	dbhandlers "javacode-test/internal/db-handlers"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

var getDbBalance = dbhandlers.DbBalance

func GetBalance(c echo.Context) error {
	cc := c.(*structs.CustomContext)
	log := ctxvalue.GetLog(cc.Ctx)

	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong UUID")
	}

	if err := uuid.Validate(id.String()); err != nil {
		return c.JSON(http.StatusBadRequest, "wrong UUID")
	}

	balance, err := getDbBalance(cc.Ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.JSON(http.StatusBadRequest, "wrong UUID")
		}
		log.Error("get balance from db", "error", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	responseWallet := structs.Wallet{ID: id, Operation: "BALANCE", Balance: balance}

	return c.JSON(http.StatusOK, responseWallet)
}

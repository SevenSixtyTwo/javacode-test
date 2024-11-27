package apihandlers

import (
	"context"
	"io"
	structs "javacode-test/api/structs"
	ctxvalue "javacode-test/internal/ctx-value"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	id := uuid.New()
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	balance := 100.22

	getDbBalance = func(ctx context.Context, id uuid.UUID) (float64, error) {
		return balance, nil
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/wallets/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(id.String())

	ctx := context.WithValue(context.Background(), ctxvalue.ValueLog, log)
	cc := &structs.CustomContext{c, ctx}
	c.Set("customContext", cc)

	if assert.NoError(t, GetBalance(cc)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, `{"valletid":`+"\""+id.String()+"\","+`"operationType":"BALANCE","amount":`+fmt.Sprintf("%.2f", balance)+`}n`, rec.Body.String())
	}

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/api/v1/wallets/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("notauuid")
	cc = &structs.CustomContext{c, ctx}
	c.Set("customContext", cc)

	if assert.NoError(t, GetBalance(cc)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		// assert.Equal(t, "wrong UUID", rec.Body.String())
	}
}

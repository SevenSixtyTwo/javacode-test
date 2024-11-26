package structs

import (
	"context"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Ctx context.Context
}

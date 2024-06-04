package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (v *DebuggerRoute) HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello World")
}

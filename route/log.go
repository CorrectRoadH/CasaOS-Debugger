package route

import (
	"net/http"

	"github.com/CorrectRoadH/CasaOS-Debugger/codegen"
	"github.com/CorrectRoadH/CasaOS-Debugger/service"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (r DebuggerRoute) QueryLog(c echo.Context, params codegen.QueryLogParams) error {
	log, err := service.MyService.LogService().QueryLog(c.Request().Context(), params.Service)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, codegen.ResponseQueryLogOk{
		Message: lo.ToPtr("OK"),
		Data:    lo.ToPtr([]string{log}),
	})
}

package route

import (
	"net/http"

	"github.com/CorrectRoadH/CasaOS-Debugger/codegen"
	"github.com/CorrectRoadH/CasaOS-Debugger/codegen/message_bus"
	"github.com/CorrectRoadH/CasaOS-Debugger/service"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (r DebuggerRoute) GetAllEventType(c echo.Context, params codegen.GetAllEventTypeParams) error {
	ets := service.MyService.MessageBus().EventTypeList(c.Request().Context())

	return c.JSON(http.StatusOK, codegen.ResponseGetEventTypeListOk{
		Message: lo.ToPtr("OK"),
		Data: lo.ToPtr(lo.Filter(ets, func(et message_bus.EventType, _ int) bool {
			return et.SourceID == params.SourceId
		}),
		),
	})
}

func (r DebuggerRoute) GetAllMessages(c echo.Context, params codegen.GetAllMessagesParams) error {
	message, err := service.MyService.MessageBus().MessageHistory(params.SourceId, params.EventType, params.Offset, params.Length)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lo.ToPtr(err.Error()))
	}
	return c.JSON(http.StatusOK, codegen.ResponseQueryMessageOk{
		Message: lo.ToPtr("OK"),
		Data:    lo.ToPtr(message),
	})
}

func (r DebuggerRoute) GetAllSources(c echo.Context) error {
	sourceList, err := service.MyService.MessageBus().Sources()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lo.ToPtr(err.Error()))
	}
	return c.JSON(http.StatusOK, codegen.ResponseGetSourceListOk{
		Message: lo.ToPtr("OK"),
		Data:    lo.ToPtr(sourceList),
	})
}

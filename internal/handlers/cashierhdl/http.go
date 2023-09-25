package cashierhdl

import (
	"cashier-system/internal/core/domain"
	"cashier-system/internal/core/port"
	"cashier-system/pkg/appmsg"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HTTPHandler struct {
	cashierService port.CashierService
}

func NewHTTPHandler(c port.CashierService) *HTTPHandler {
	return &HTTPHandler{
		cashierService: c,
	}
}

func (hdl *HTTPHandler) DoOrder(c echo.Context) error {

	request := domain.DoOrderRequest{}

	err := c.Bind(&request)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(400, appmsg.BAD_REQUEST)
	}

	orderResult, err := hdl.cashierService.DoOrder(request)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(500, appmsg.INTERNAL_SERVER_ERROR)
	}

	return c.JSON(200, orderResult)

}

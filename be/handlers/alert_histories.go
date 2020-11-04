package handlers

import "github.com/labstack/echo"

func (i *impl) GetAlertHistories(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()

	var (
		limit  = 10
		offset = 0
	)

	return i.service.ListAlertHistories(ctx, limit, offset)
}

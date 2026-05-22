package http

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mihailtudos/yumgo/common"
	"github.com/mihailtudos/yumgo/common/log"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger = common.NewEchoSlogAdapter(slog.Default())

	useMiddlewares(e)
	e.HTTPErrorHandler = HandleError
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	return e
}

func HandleError(err error, c echo.Context) {
	httpCode := http.StatusInternalServerError
	msg := any("Internal server error")
	log.FromContext(c.Request().Context()).With("error", err).Error("HTTP error")

	httpErr := &echo.HTTPError{}
	if errors.As(err, &httpErr) {
		httpCode = httpErr.Code
		msg = httpErr.Message
	}

	jsonErr := c.JSON(
		httpCode,
		map[string]any{
			"error": msg,
		},
	)

	if jsonErr != nil {
		panic(err)
	}
}

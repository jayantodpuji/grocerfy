package delivery

import "github.com/labstack/echo/v4"

type Response struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	ErrorTrace any    `json:"errorTrace,omitempty"`
}

func ResponseError(c echo.Context, code int, msg string) error {
	return c.JSON(code, Response{
		Success: false,
		Message: msg,
	})
}

func ResponseSuccess(c echo.Context, code int, data any) error {
	return c.JSON(code, Response{
		Success: true,
		Data:    data,
	})
}

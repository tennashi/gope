package handler

import (
	"github.com/labstack/echo"
	"github.com/tennashi/gope/context"
)

// Func represent handler function with middleware.Context
type Func func(c *context.Context) error

// C convert Func into echo.HandlerFunc
func C(h Func) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*context.Context))
	}
}

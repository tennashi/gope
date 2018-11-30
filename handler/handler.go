package handler

import (
	"github.com/labstack/echo"
	"github.com/tennashi/gope/context"
)

// Func represent a handler function that takes context.Project as an argument.
type Func func(c *context.Project) error

// PC convert Func to echo.HandlerFunc.
func PC(h Func) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*context.Project))
	}
}

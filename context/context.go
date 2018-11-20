package context

import (
	"github.com/labstack/echo"
	"github.com/tennashi/gope/config"
)

// Context is custom context
type Context struct {
	echo.Context
	Config *config.Config
}

// NewContext construct Context
func NewContext(ec echo.Context, config *config.Config) *Context {
	return &Context{ec, config}
}

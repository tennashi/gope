package router

import (
	"github.com/labstack/echo"
	"github.com/tennashi/gope/config"
	"github.com/tennashi/gope/context"
	"github.com/tennashi/gope/handler"
)

// Init initialize echo.Echo
func Init(config *config.Config) *echo.Echo {
	e := echo.New()

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			ctx := context.NewContext(ec, config)
			return h(ctx)
		}
	})

	v1 := e.Group("/api/v1")
	{
		hh := handler.NewHost()
		v1.GET("/hosts/:name", handler.C(hh.GetHosts))
		//v1.GET("/procedure", handler.C(handler.GetProcedures))
	}

	return e
}

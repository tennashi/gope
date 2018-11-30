package router

import (
	"github.com/labstack/echo"
	"github.com/tennashi/gope/handler"
	"github.com/tennashi/gope/store"
)

// Init initialize echo.Echo
func Init(projStore *store.Project) *echo.Echo {
	e := echo.New()

	v1 := e.Group("/api/v1")
	{
		pjh := handler.NewProject(projStore)
		p := v1.Group("/:project_name", pjh.UseProjectContext)
		{
			p.GET("", handler.PC(pjh.GetProject))
			hh := handler.NewHost()
			{
				p.GET("/hosts", handler.PC(hh.GetHostFiles))
				p.GET("/hosts/:name", handler.PC(hh.GetHosts))
				p.GET("/hosts/:name/:host", handler.PC(hh.GetHost))
			}
			ph := handler.NewProcedure()
			{
				p.GET("/procedures", handler.PC(ph.GetProcedureFiles))
				p.GET("/procedures/:name", handler.PC(ph.GetProcedure))
			}

		}
	}

	return e
}

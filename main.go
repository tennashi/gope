package main

import (
	"github.com/tennashi/gope/config"
	"github.com/tennashi/gope/router"
)

func main() {
	config := config.NewConfig("./")
	e := router.Init(config)
	e.Logger.Error(e.Start(":" + config.GetString("port")))
}

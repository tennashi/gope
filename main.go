package main

import (
	"log"

	"github.com/tennashi/gope/config"
	"github.com/tennashi/gope/router"
	"github.com/tennashi/gope/store"
)

func main() {
	flag := config.InitFlags()
	baseConf := config.NewBaseConfig(flag.ConfigFile)

	projStore, err := store.NewProject(baseConf.GetString("project_file"))
	if err != nil {
		log.Println(err)
		return
	}

	e := router.Init(projStore)
	e.Logger.Error(e.Start(":" + baseConf.GetString("port")))
}

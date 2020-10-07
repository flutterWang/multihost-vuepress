package main

import (
	"github.com/mygomod/muses"
	"github.com/mygomod/muses/pkg/cmd"
	musgin "github.com/mygomod/muses/pkg/server/gin"
	"github.com/mygomod/muses/pkg/server/stat"
	"multihost-vuepress/app/pkg/conf"
	"multihost-vuepress/app/pkg/mus"
	"multihost-vuepress/app/router"
)

func main() {
	app := muses.Container(
		cmd.Register,
		stat.Register,
		musgin.Register,
	)
	app.SetGinRouter(router.InitRouter)
	app.SetPostRun(mus.Init,conf.Init)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}

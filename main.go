package main

import (
	"go_test/config"
	"go_test/route"
)

const (
	//mongodb://localhost:27017
	mongoDBEnPint = ""
	portWebServie = ":8000"
)

func main() {
	db := config.NewDatabaseMgo(mongoDBEnPint)

	app := config.NewframeworkFiber()
	app.Default()
	route.SetupRouters(app, db)
	app.Run(portWebServie)
}

package main

import (
	"github.com/nelsonomoi/usingJWT/config"
	"github.com/nelsonomoi/usingJWT/routes"
)

func main() {

	// Database Initialization
	config.Connect("root:@tcp(localhost:3306)/usingjwt?parseTime=true")
	config.Migrate()

	router := routes.InitRoutes()

	router.Run(":3000")

}

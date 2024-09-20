package main

import (
	"Gim/internal/logic"
	"Gim/internal/router"
	"Gim/internal/server"
)

func main() {
	// testGin()
	// testGorm()

	server.InitMySQL()
	logic.InitUserTable()
	server.InitRedis()

	r := router.Router()
	r.Run()
}

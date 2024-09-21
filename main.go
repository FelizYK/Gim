package main

import (
	"Gim/internal/router"
	"Gim/internal/server"
)

func main() {
	// testGin()
	// testGorm()

	server.InitMySQL()
	// logic.InitUserTable()
	// logic.InitMessageTable()
	// logic.InitRelationTable()
	// logic.InitGroupTable()

	server.InitRedis()

	r := router.Router()
	r.Run()
}

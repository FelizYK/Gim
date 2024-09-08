package main

import (
	"Gim/internal/router"
	"Gim/internal/sql"
	// "Gim/utils"
)

func main() {
	// testGin()
	// testGorm()

	// utils.InitConfig()
	sql.InitMySQL()

	r := router.Router()
	r.Run()
}

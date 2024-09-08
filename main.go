package main

import (
	"Gim/internal/logic"
	"Gim/internal/router"
	"Gim/internal/sql"
	// "Gim/utils"
)

func main() {
	// testGin()
	// testGorm()

	// utils.InitConfig()
	sql.InitMySQL()
	logic.InitUserTable()

	r := router.Router()
	r.Run()
}

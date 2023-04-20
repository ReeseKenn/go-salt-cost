package cmd

import (
	"salt-cost/internal/database"
	"salt-cost/internal/router"
)

func RunServer() {
	database.Connect()
	database.CreateTables()
	defer database.Close()
	r := router.New()
	r.Run(":8080")
}

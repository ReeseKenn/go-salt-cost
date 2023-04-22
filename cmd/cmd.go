package cmd

import (
	"salt-cost/internal/database"
	"salt-cost/internal/router"
)

func RunServer() {
	database.MysqlConnect()
	database.MysqlCreateTable()
	defer database.MysqlClose()
	r := router.New()
	r.Run(":8080")
}

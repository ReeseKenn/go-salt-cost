package cmd

import (
	"salt-cost/internal/router"
)

func RunServer() {
	r := router.New()
	r.Run(":8080")
}

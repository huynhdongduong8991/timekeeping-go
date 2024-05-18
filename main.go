package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"time-keeping.com/jobs"
	"time-keeping.com/lib"
	"time-keeping.com/routes"
)

func main() {
	// Setup database
	err := lib.ConnectDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Listen to events
	go jobs.ListenToEvents()

	// Setup gin server
	router := gin.Default()
	routes.Router(router)
	router.Run(":3000")
}

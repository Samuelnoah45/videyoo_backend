package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"server/config"
	log "server/logs"

	// db "server/pkgs/DB"
	auth_router "server/pkgs/auth/router"
)

func main() {
	// database, err := db.Connect()
	// if err != nil {
	// 	panic(err)
	// }
	// defer database.Close()

	// fmt.Println("---------------------#########Successfully connected to the database!##########---------")
	log.InfoLevel(fmt.Sprintf("central management server running on port %s", config.PORT))
	router := gin.Default()
	// Auth routes
	auth_router.SetupRoutes(router)

	// File upload router
	router.Run(fmt.Sprintf(":%s", config.PORT))
}

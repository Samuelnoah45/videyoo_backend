package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"server/config"
	log "server/logs"
	db "server/pkgs/DB"
	auth_router "server/pkgs/auth/router"
	fileupload "server/pkgs/file_upload"
	request_router "server/pkgs/request/router"
	transaction_router "server/pkgs/transaction/router"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	fmt.Println("---------------------#########Successfully connected to the database!##########---------")
	log.InfoLevel(fmt.Sprintf("central management server running on port %s", config.PORT))
	router := gin.Default()
	// Auth routes
	auth_router.SetupRoutes(router)
	// Request routes
	request_router.SetupRoutes(router)
	// Transaction routes

	transaction_router.SetupRoutes(router)

	router.POST("/file_upload", fileupload.UploadImage)
	// File upload router
	router.Run(fmt.Sprintf(":%s", config.PORT))
}

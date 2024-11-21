package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"samb-backend/config"
	"samb-backend/routes"
)

func main(){
	database := config.InitDB()
	defer func() {
		sqlDB, err := database.DB()
		if err != nil {
			log.Fatalf("Error accessing underlying sql.DB: %v", err)
		}
		sqlDB.Close() 
	}()

	router := routes.RegisteredRoutes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		os.Exit(0)
	}()

	log.Println("Development server is on port: 8080...")
	router.Run(":8080")
}
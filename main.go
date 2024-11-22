package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"samb-backend/config"
	"samb-backend/routes"

	"github.com/gorilla/handlers"
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

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000", "https://magenta-smakager-3db320.netlify.app"}) // Allow the frontend URL
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	handleRouterCORS :=handlers.CORS(originsOk, headersOk, methodsOk)(router)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		os.Exit(0)
	}()

	log.Println("Development server is on port: 8080...")
	if err := http.ListenAndServe(":8080",handleRouterCORS); err!= nil {
		log.Fatalf("Development server failed: %v",err)
	}
}
package app

import (
	"blockchain_to_go/controllers"
	"blockchain_to_go/routers"
	"log"
	"net/http"
)

func Run() {
	app := http.NewServeMux()

	chain := controllers.InitBlockchain()
	router := routers.InitRoutes(chain)

	// Add /version01 prefix to all routes
	app.Handle("/version01/", http.StripPrefix("/version01", router))

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app))
}

package main

import (
	"blockchain_to_go/controllers"
	"blockchain_to_go/routers"
	"log"
	"net/http"
)

func main() {
	chain := controllers.InitBlockchain()
	router := routers.InitRoutes(chain)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

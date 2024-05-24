package routers

import (
	"blockchain_to_go/controllers"
	"net/http"
)

// InitRoutes initializes the routes for the HTTP server
func InitRoutes(chain *controllers.Blockchain) *http.ServeMux {
	mux := http.NewServeMux()

	// Define the handlers for blockchain operations
	mux.HandleFunc("/get", controllers.GetBlockchainHandler)
	mux.HandleFunc("/add", controllers.AddBlockHandler)

	return mux
}

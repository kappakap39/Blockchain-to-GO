package main

import (
	"GO-Project-blockchain-API02/controllers"
	"GO-Project-blockchain-API02/routers"
	"fmt"
	"net/http"
)

func main() {
	chain := controllers.InitBlockchain()

	http.HandleFunc("/blocks", routers.GetBlockchainHandler(chain))
	http.HandleFunc("/addblock", routers.AddBlockHandler(chain))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

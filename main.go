package main

import (
	"blockchain_to_go/controllers"
	"blockchain_to_go/routers"
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

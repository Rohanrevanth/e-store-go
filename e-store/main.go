package main

import (
	"github.com/Rohanrevanth/e-store-go/database"
	"github.com/Rohanrevanth/e-store-go/http"
)

func main() {
	database.ConnectDatabase()
	// database.InitializeRedis()
	http.StartServer()
}

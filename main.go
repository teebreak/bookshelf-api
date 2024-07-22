package main

import (
	"bookshelf-api/database"
	"bookshelf-api/routers"
	"log"
	"net/http"
)

func main() {
	database.Connect()
	database.InitRedis()
	database.InitElasticsearch()

	r := routers.Router()
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

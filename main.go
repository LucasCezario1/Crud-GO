package main

import (
	"filmes-crud/database"
	server "filmes-crud/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}

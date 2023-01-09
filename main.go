package main

import (
	"GOLANG/db"
	"GOLANG/server"
	"log"
)

func main() {
	err := db.Connect()

	if err != nil {
		log.Println("Failed to connect to DB")
	}

	server.Init()
}

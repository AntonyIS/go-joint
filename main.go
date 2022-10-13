package main

import (
	"log"

	"github.com/AntonyIS/go-joint/api"
)

func main() {
	// Get Attendee repository
	router, err := api.AppRouter()
	if err != nil {
		log.Fatal(err)
	}
	router.Run(":8080")
}

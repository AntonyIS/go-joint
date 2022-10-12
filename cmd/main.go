package main

import (
	"log"
	"net/http"

	"github.com/AntonyIS/go-joint/api"
	"github.com/AntonyIS/go-joint/app"
	"github.com/AntonyIS/go-joint/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get Attendee repository
	repo, err := repository.NewAttendeeRepository()
	if err != nil {
		log.Fatal(err)
	}
	// Get attendee service
	srv := app.AttendeeService(repo)
	// Get Attendee handler
	attendeeHandler := api.NewAttendeeHandler(srv)
	// Gin router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Welcome to Go joint application"}) })
	// Attendee Handler
	router.GET("/attendees", attendeeHandler.GetAll)
	router.GET("/attendees/:id", attendeeHandler.Get)
	router.POST("/attendees/", attendeeHandler.Post)
	router.PUT("/attendees/:id", attendeeHandler.Put)
	router.DELETE("/attendees/:id", attendeeHandler.Delete)

	router.Run(":8080")

}

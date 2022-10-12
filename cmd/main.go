package main

import (
	"net/http"

	"github.com/AntonyIS/go-joint/api"
	"github.com/AntonyIS/go-joint/app"
	"github.com/AntonyIS/go-joint/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get Attendee repository
	repo := repository.NewAttendeeRepository()
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
	router.GET("/attendees/", attendeeHandler.Post)
	router.GET("/attendees/:id", attendeeHandler.Put)
	router.GET("/attendees/:id", attendeeHandler.Delete)

	router.Run(":8080")

}

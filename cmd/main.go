package main

import (
	"net/http"

	h "go-joint/api"
	"go-joint/app"
	repo "go-joint/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get Attendee repository
	repo := repo.NewAttendeeRepository()
	// Get attendee service
	srv := app.AttendeeService(repo)
	// Get Attendee handler
	attendeeHandler := h.NewAttendeeHandler(srv)
	// Gin router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Welcome to Go joint application"}) })
	// Attendee Handler
	router.GET("/attendees", attendeeHandler.GetAll)
	router.GET("/attendees/:id", attendeeHandler.Get)
	router.GET("/attendees/", attendeeHandler.Create)
	router.GET("/attendees/:id", attendeeHandler.Update)
	router.GET("/attendees/:id", attendeeHandler.Delete)

	router.Run(":8080")

}

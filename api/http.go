package api

import (
	"net/http"

	"github.com/AntonyIS/go-joint/app"
	"github.com/AntonyIS/go-joint/repository"
	"github.com/gin-gonic/gin"
)

type RouteHandler interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	Post(*gin.Context)
	Put(*gin.Context)
	Delete(*gin.Context)
}

func AppRouter() (*gin.Engine, error) {
	repo, err := repository.NewAttendeeRepository()
	if err != nil {
		return nil, err
	}
	// Get attendee service
	srv := app.AttendeeService(repo)
	// Get Attendee handler
	attendeeHandler := NewAttendeeHandler(srv)
	// Gin router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Welcome to Go joint application"}) })
	// Attendee Handler
	router.GET("/attendees", attendeeHandler.GetAll)
	router.GET("/attendees/:id", attendeeHandler.Get)
	router.POST("/attendees/", attendeeHandler.Post)
	router.PUT("/attendees/:id", attendeeHandler.Put)
	router.DELETE("/attendees/:id", attendeeHandler.Delete)

	return router, nil

}

package api

import (
	"go-joint/app"

	"github.com/gin-gonic/gin"
)

type RouteHandler interface {
	Get(*gin.Context)
	POST(*gin.Context)
	PUT(*gin.Context)
	DELETE(*gin.Context)
}

type AttendeeHandler struct {
	attendeeService app.AttendeeService
}
type EventHandler struct {
	eventService app.EventService
}
type TicketHandler struct {
	ticketService app.TicketService
}

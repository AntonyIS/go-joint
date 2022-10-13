package api

import (
	"net/http"

	"github.com/AntonyIS/go-joint/app"

	"github.com/gin-gonic/gin"
)

type attendeeHandler struct {
	attendeeService app.AttendeeService
}

func NewAttendeeHandler(attendeeService app.AttendeeService) RouteHandler {
	return attendeeHandler{
		attendeeService,
	}
}

func (h attendeeHandler) Post(c *gin.Context) {
	var attendee app.Attendee

	if err := c.ShouldBindJSON(&attendee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	_, err := h.attendeeService.Create(&attendee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"attendee": attendee,
	})

}

func (h attendeeHandler) Get(c *gin.Context) {
	id := c.Param("id")
	attendee, err := h.attendeeService.Read(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"attendee": attendee})
}

func (h attendeeHandler) GetAll(c *gin.Context) {
	attendees, err := h.attendeeService.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"attendees": attendees})
}

func (h attendeeHandler) Put(c *gin.Context) {
	var attendee app.Attendee

	if err := c.ShouldBindJSON(&attendee); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	data, err := h.attendeeService.Update(&attendee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"attendee": data,
	})
}

func (h attendeeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.attendeeService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "attendee deleted successfuly",
	})
}

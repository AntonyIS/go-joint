package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AntonyIS/go-joint/api"
	"github.com/AntonyIS/go-joint/app"
	"github.com/AntonyIS/go-joint/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func RouterEngine() *gin.Engine {
	return gin.Default()
}

func AttendeeHandler() api.RouteHandler {
	repo, err := repository.NewAttendeeRepository()
	if err != nil {
		return nil
	}
	// Get attendee service
	srv := app.AttendeeService(repo)
	// Get Attendee handler
	attendeeHandler := api.NewAttendeeHandler(srv)
	// Gin router
	return attendeeHandler
}

func TestHomeRoute(t *testing.T) {
	mockResponse := `{"message": "Welcome to Go joint application"}`
	handler := AttendeeHandler()

	r := RouterEngine()
	r.GET("/", handler.Get)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	response, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(response))
	assert.Equal(t, http.StatusOK, w.Code)
}

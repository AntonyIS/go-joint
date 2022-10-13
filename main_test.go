package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func RouterEngine() *gin.Engine {
	return gin.Default()
}

func TestHomeRoute(t *testing.T) {
	mockResponse := `{"message": "Welcome to Go joint application"}`

	r := RouterEngine()
	r.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Welcome to Go joint application"}) })
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	response, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(response))
	assert.Equal(t, http.StatusOK, w.Code)
}

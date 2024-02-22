package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/furqanalimir/commuter/handlers"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetFruitHandler(t *testing.T) {
	mockResponse := `{"data": {"id": 1,"name":"Water Mellon","price": 1.1,"currency": "usd","Origin":"africa"},"error":null}`
	router := SetUpRouter()
	router.GET("/api/v0.1/fruits/:id", handlers.HandlerGetFruit)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v0.1/fruits/1", nil)
	// req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("admin:password")))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	log.Println("Body::", w.Body.String())
	assert.Equal(t, mockResponse, w.Body.String())
}

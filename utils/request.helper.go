package helper

import (
	"log"

	"github.com/gin-gonic/gin"
)

type SwaggerRequestResponse struct {
	Error string
	Data  string
}

func ReqResHelper(c *gin.Context, s int, data, err any) {
	if err != nil {
		log.Println("Error: ", err)
	}
	c.JSON(s, gin.H{
		"error": err,
		"data":  data,
	})
}

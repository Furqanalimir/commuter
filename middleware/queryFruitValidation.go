package middleware

import (
	"log"
	"net/http"
	"strconv"

	helper "github.com/furqanalimir/commuter/utils"

	"github.com/gin-gonic/gin"
)

func QueryValidationMiddleware(c *gin.Context) {
	log.Println("query: ", c.Query("id"))
	id, err := strconv.Atoi(c.Param("id"))
	log.Println("id: ", id)
	if err != nil {
		helper.ReqResHelper(c, http.StatusBadRequest, nil, "Id is required and must be an integer")
		c.Abort()
		return
	}
	c.Set("id", id)
	c.Next()
}

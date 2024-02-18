package validator

import (
	helper "microservices/micro-service/commuter/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryValidationMiddleware(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ReqResHelper(c, http.StatusBadRequest, nil, "Id is required and must be an integer")
		c.Abort()
		return
	}
	c.Set("id", id)
	c.Next()
}

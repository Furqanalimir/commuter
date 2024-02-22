package middleware

import (
	"net/http"
	"strconv"

	helper "github.com/furqanalimir/commuter/utils"

	"github.com/gin-gonic/gin"
)

// validate id is passed in params or not
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

// validate id is passed in params or not
func QueryValidationMiddlewareD(field string, errMsg string) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param(field))
		if err != nil {
			helper.ReqResHelper(c, http.StatusBadRequest, nil, errMsg)
			c.Abort()
			return
		}
		c.Set("id", id)
		c.Next()
	}
}

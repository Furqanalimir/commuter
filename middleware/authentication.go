package middleware

import (
	"net/http"

	helper "github.com/furqanalimir/commuter/utils"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	// validate token
	err := helper.TokenValid(c.Request)
	if err != nil {
		helper.ReqResHelper(c, http.StatusNonAuthoritativeInfo, nil, err.Error())
		c.Abort()
	}
	c.Next()
}

func AuthenticateAdmin(c *gin.Context) {
	// verify if user is admin or not
	err := helper.VerifyAdmin(c.Request)
	if err != nil {
		helper.ReqResHelper(c, http.StatusNonAuthoritativeInfo, nil, err.Error())
		c.Abort()
	}
	c.Next()
}

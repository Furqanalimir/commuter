package middleware

import (
	"net/http"

	helper "github.com/furqanalimir/commuter/utils"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	err := helper.TokenValid(c.Request)
	if err != nil {
		helper.ReqResHelper(c, http.StatusNonAuthoritativeInfo, nil, err.Error())
		c.Abort()
	}
	c.Next()
}

func AuthenticateAdmin(c *gin.Context) {
	err := helper.VerifyAdmin(c.Request)
	if err != nil {
		helper.ReqResHelper(c, http.StatusNonAuthoritativeInfo, nil, err.Error())
		c.Abort()
	}
	c.Next()
}

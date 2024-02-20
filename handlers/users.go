package handlers

import (
	"microservices/micro-service/commuter/data"
	validator "microservices/micro-service/commuter/middleware"
	helper "microservices/micro-service/commuter/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserConfig struct {
	R        *gin.Engine
	BasePath string
}

func NewUserHandler(c *UserConfig) {
	// Create an fruits group
	g := c.R.Group(c.BasePath + "/users")

	g.POST("/", handlerAddUser)
	g.POST("/login", handleLogin)

	// routes with Query middleware validation
	g.GET("/verify", handlerVerifyUser)
	g.Use(validator.QueryValidationMiddleware)
	// g.Use(helper.ExtractTokenAuth)
	// g.PUT("/", handlerUpdateUser)
	// g.DELETE("/:id", handlerRemoveFruit)
}

func handlerAddUser(c *gin.Context) {
	u := &data.User{}
	err := u.ToJSON(c.Request.Body)
	if err != nil {
		helper.ReqResHelper(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	// custom_role_validation
	err = u.Validate()
	if err != nil {
		helper.ReqResHelper(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	err = data.AddUser(u)
	if err != nil {
		helper.ReqResHelper(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	helper.ReqResHelper(c, http.StatusOK, true, nil)
}

func handleLogin(c *gin.Context) {
	var u data.User
	if err := c.ShouldBindJSON(&u); err != nil {
		helper.ReqResHelper(c, http.StatusUnprocessableEntity, nil, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	id, err := u.VerifyUser()
	if err != nil {
		helper.ReqResHelper(c, http.StatusUnprocessableEntity, nil, err.Error())
		return
	}
	token, _ := helper.CreateToken(id)
	helper.ReqResHelper(c, http.StatusOK, token, nil)
}

func handlerVerifyUser(c *gin.Context) {
	validity, err := helper.GetTokenValidity(c.Request)
	// tokenAuth
	if err != nil {
		helper.ReqResHelper(c, http.StatusUnauthorized, nil, "unauthorized, please login to get access to protected info.")
		return
	}
	helper.ReqResHelper(c, http.StatusOK, validity.UTC().String(), nil)
}

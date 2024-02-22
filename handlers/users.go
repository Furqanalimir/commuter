package handlers

import (
	"net/http"

	"github.com/furqanalimir/commuter/data"
	validator "github.com/furqanalimir/commuter/middleware"
	helper "github.com/furqanalimir/commuter/utils"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
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

// AddUser		godoc
// @Summary		Create user
// @Description	Save User data
// @Param		user body data.User true "create user"
// @produce		applicaton/json
// @Tags		user
// @Success		200	{object} gin.H  "create response"
// @Success		400	{object} gin.H  "error response"
// @Router		/users [post]
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

// Login		godoc
// @Summary		Create authentication token
// @Description	validate user and generate token
// @Param       Authentication body data.Authentication true "user email and password"
// @produce		applicaton/json
// @Tags		user
// @Success		200	{object} gin.H "token"
// @Success		422	{object} gin.H "invalid data"
// @Router		/users/login [post]
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

// Verify		godoc
// @Summary		Verify user token
// @Description	Authenticate user token
// @produce		applicaton/json
// @Tags		user
// @Success		200	{object} gin.H "time stamp"
// @Success		401	{object} gin.H "unauthorized message"
// @Router		/users/verify [get]
// @Security ApiKeyAuth
func handlerVerifyUser(c *gin.Context) {
	validity, err := helper.GetTokenValidity(c.Request)
	// tokenAuth
	if err != nil {
		helper.ReqResHelper(c, http.StatusUnauthorized, nil, "unauthorized, please login to get access to protected info.")
		return
	}
	helper.ReqResHelper(c, http.StatusOK, validity.UTC().String(), nil)
}

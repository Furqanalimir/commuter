package handlers

import (
	"microservices/micro-service/commuter/data"
	validator "microservices/micro-service/commuter/middleware"
	helper "microservices/micro-service/commuter/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FruitsConfig struct {
	R *gin.Engine
}

func NewHandler(c *FruitsConfig) {
	// Create an fruits group
	g := c.R.Group("/fruits")

	g.GET("/", handlerGetAllFruits)
	g.POST("/", handlerAddFurit)

	// routes with Query middleware validation
	g.Use(validator.QueryValidationMiddleware)
	g.GET("/:id", handlerGetFruit)
	g.DELETE("/:id", handlerRemoveFruit)
}

func handlerGetFruit(c *gin.Context) {
	id := c.MustGet("id").(int)
	f, err := data.GetFruit(id)
	if err != nil {
		helper.ReqResHelper(c, http.StatusOK, nil, err.Error())
		return
	}
	helper.ReqResHelper(c, http.StatusOK, f, nil)
}

func handlerAddFurit(c *gin.Context) {
	fruit := &data.Fruit{}
	fruit.ToJSON(c.Request.Body)
	err := fruit.Validate()
	if err != nil {
		helper.ReqResHelper(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	fl, err := data.AddFruit(fruit)
	if err != nil {
		helper.ReqResHelper(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	helper.ReqResHelper(c, http.StatusOK, fl, nil)
}

func handlerGetAllFruits(c *gin.Context) {
	f := data.GetAllFuits()
	helper.ReqResHelper(c, http.StatusOK, f, nil)
}

func handlerRemoveFruit(c *gin.Context) {
	id := c.MustGet("id").(int)
	err := data.RemoveFruit(id)
	if err != nil {
		helper.ReqResHelper(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	helper.ReqResHelper(c, http.StatusOK, true, nil)
}

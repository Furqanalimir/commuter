// Package Classification of Fruits API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// swagger:meta
package handlers

import (
	"microservices/micro-service/commuter/data"
	middleware "microservices/micro-service/commuter/middleware"
	helper "microservices/micro-service/commuter/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FruitConfig struct {
	R        *gin.Engine
	BasePath string
}

func NewFruitHandler(c *FruitConfig) {
	// Create an fruits group
	g := c.R.Group(c.BasePath + "/fruits")

	g.GET("/", handlerGetAllFruits)
	g.Use(middleware.Authentication)
	g.Use(middleware.AuthenticateAdmin)
	g.POST("/", handlerAddFurit)

	// routes with Query middleware validation
	g.Use(middleware.QueryValidationMiddleware)
	// g.Use(helpers.)
	g.GET("/:id", handlerGetFruit)
	g.DELETE("/:id", handlerRemoveFruit)
}

func handlerGetFruit(c *gin.Context) {
	id := c.MustGet("id").(int)
	f, err := data.GetFruit(id)
	if err != nil {
		helper.ReqResHelper(c, http.StatusNotFound, nil, err.Error())
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

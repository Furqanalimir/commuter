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
	"net/http"

	"github.com/furqanalimir/commuter/data"
	"github.com/furqanalimir/commuter/middleware"
	helper "github.com/furqanalimir/commuter/utils"
	"github.com/gin-gonic/gin"
)

type FruitConfig struct {
	R        *gin.Engine
	BasePath string
}

func NewFruitHandler(c *FruitConfig) {
	// Create an fruits group
	g := c.R.Group(c.BasePath + "/fruits")

	g.Use(middleware.Authentication)
	g.GET("/", handlerGetAllFruits)
	g.GET("/:id", middleware.QueryValidationMiddleware, handlerGetFruit)

	g.Use(middleware.AuthenticateAdmin)
	g.POST("/", handlerAddFurit)
	g.Use(middleware.QueryValidationMiddleware)
	// routes with Query middleware validation
	g.DELETE("/:id", handlerRemoveFruit)
}

// Fetch Fruit		godoc
// @Summary		fruits
// @Description	get fruit info by id
// @produce		applicaton/json
// @Success		200	{object} gin.H "time stamp"
// @Success		401	{object} gin.H "unauthorized message"
// @Param		id path int true "get fruit by id"
// @Router		/fruits/{id} [get]
// @Security ApiKeyAuth
func handlerGetFruit(c *gin.Context) {
	id := c.GetInt("id")
	// .(int)
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

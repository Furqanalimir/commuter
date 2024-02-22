package main

import (
	"context"
	"log"

	// _ "swag-gin-demo/docs"
	_ "github.com/furqanalimir/commuter/docs"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/furqanalimir/commuter/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	basePath = "/api/v0.1"
)

// @BasePath /api/v0.1
// @title Go + Gin User API
// @version 1.0
// @description This is a sample server user server. You can visit the GitHub repository at https://github.com/Furqanalimir/commuter

// @contact.name API Support
// @contact.url https://furqanali.vercel.app/
// @contact.email mirfurqan89@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// url := ginSwagger.URL("swagger/doc.json")
	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handlers.NewFruitHandler(&handlers.FruitConfig{
		R:        router,
		BasePath: basePath,
	})
	handlers.NewUserHandler(&handlers.UserConfig{
		R:        router,
		BasePath: basePath,
	})
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 5,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()
	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}

package main

import (
	"context"
	"fmt"

	"net/http"
	"one_pay/config"
	"one_pay/routes"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	server *http.Server
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)

	r := gin.Default()

	// initialize out configuration
	config.InitCofig()
	serverConfig := config.AppConfig.Server
	routes.InitializeRoutes(r, &config.AppConfig)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Initialize server
	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", serverConfig.Port),
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	gracefulShutdown()
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	<-quit

	log.Println("----shutdown server----")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("----shutdown failed-----%+v", err)
	}

	log.Println("exiting server")
}

package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"site-portfolio/config"
)

func SetupGracefulShutdown(router *gin.Engine) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("\nShutting down server...")
		config.CloseDB()
		os.Exit(0)
	}()
}

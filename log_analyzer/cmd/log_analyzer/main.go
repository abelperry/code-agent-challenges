package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// TODO: start server

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// TODO: stop server
}

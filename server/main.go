package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "aeroline/docs"
)

// 	   _______ __
//    / ____(_) /_  ___  _____
//   / /_  / / __ \/ _ \/ ___/
//  / __/ / / /_/ /  __/ /
// /_/   /_/_.___/\___/_/

// @title Fiber API
// @version 1.0
// @host localhost:7000
// @BasePath /api
func main() {
	app, cleanup, err := initApp()
	if err != nil {
		panic(err)
	}

	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := app.Listen(":7000"); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-sigChan:
		cleanup()
		log.Println("Server stopped gracefully")

	case err := <-errChan:
		log.Printf("Server error: %v", err)
		cleanup()
	}
}

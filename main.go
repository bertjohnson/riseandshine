package main

import (
	"context"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()

	// Run web server
	c := context.Background()
	go RunWebServer(c)

	// Block the main goroutine to keep the program running
	ch := make(chan struct{})
	<-ch
}

package main

import (
    "fmt"
    "github.com/ryan-tech/gameoflife/internal/server"
)


func main() {
	// start server
	fmt.Println("Starting server...")
	server.Start()
}

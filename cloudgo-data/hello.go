package main

import (
	"github.com/user/hello/services"
)

const (
	PORT string = "8000"
)

func main() {
	server := services.NewServer()
	server.Run(":" + PORT)
}
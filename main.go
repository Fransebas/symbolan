package main

import (
	"os"
	"symbolan/server"
)

func main() {
	port := os.Getenv("PORT")
	server.Run(port)
}

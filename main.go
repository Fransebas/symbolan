package main

import (
	"os"
	"symbolan/server"
)

func main() {
	port := os.Getenv("FOO")
	server.Run(port)
}

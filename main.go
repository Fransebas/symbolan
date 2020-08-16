package main

import (
	"os"
	"symbolan/server"
)

func main() {
	server.Run(os.Args[1])
}

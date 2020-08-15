package main

import (
	"os"
	"symbolan/main/server"
)

func main() {
	server.Run(os.Args[1])
}

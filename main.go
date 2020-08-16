package main

import (
	"fmt"
	"os"
	"symbolan/server"
)

func main() {
	fmt.Println("Hello")
	server.Run(os.Args[1])
}

package main

import (
	"flag"
	"fmt"

	"github.com/DiscoViking/rogue/client"
	"github.com/DiscoViking/rogue/server"
)

var asServer = flag.Bool("server", false, "run in server mode")
var asClient = flag.Bool("client", false, "run in client mode")

func main() {
	flag.Parse()

	if *asServer {
		fmt.Printf("Running server on port %v\n", 12345)
		server.Run("12345")
	} else if *asClient {
		fmt.Printf("Connecting client to %v\n", "localhost:12345")
		client.Run("localhost:12345")
	}
}

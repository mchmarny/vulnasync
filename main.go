package main

import "github.com/mchmarny/vulnasync/internal/server"

const (
	name = "vulnasync"
)

var (
	// Set at build time.
	version = "v0.0.1-default"
)

func main() {
	server.Run(name, version)
}

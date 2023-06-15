package main

import (
	"github.com/github/internal/server"
)

var (
	// Could be set at build time.
	version = "v0.0.1-default"
	name    = "demo"
)

func main() {
	server.Run(name, version)
}

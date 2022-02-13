package main

import (
	"jasoncoding.com/backendv2/config"
	"jasoncoding.com/backendv2/server"
)

func main() {
	config.Init()
	server.Init()
}
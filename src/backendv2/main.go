package main

import (
	"jasoncoding.com/backendv2/config"
	"jasoncoding.com/backendv2/cool"
	"jasoncoding.com/backendv2/server"
)

func main() {
	config.Init()
	cool.Init(cool.New())
	server.Init()
}

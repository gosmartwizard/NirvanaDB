package main

import (
	"flag"
	"fmt"
)

const (
	SERVER_HOST      = "localhost"
	SERVER_HOST_PORT = 4949
	SERVER_TYPE      = "tcp"
)

var serverHost string
var serverHostPort int
var serverType string

func init() {
	flag.StringVar(&serverHost, "host", SERVER_HOST, "Server IP address")
	flag.IntVar(&serverHostPort, "port", SERVER_HOST_PORT, "port for the Server")
	flag.StringVar(&serverType, "type", SERVER_TYPE, "Network type for the Server")
}

func main() {

	flag.Parse()

	fmt.Println(serverHost, serverHostPort, serverType)
}

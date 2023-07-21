package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "4949"
	SERVER_TYPE = "tcp"
)

var serverHost string
var serverPort string
var serverType string

func init() {
	flag.StringVar(&serverHost, "host", SERVER_HOST, "Server IP address")
	flag.StringVar(&serverPort, "port", SERVER_PORT, "port for the Server")
	flag.StringVar(&serverType, "type", SERVER_TYPE, "Network type for the Server")
}

func main() {

	flag.Parse()

	fmt.Println(serverHost, serverPort, serverType)

	startServer()
}

func startServer() {
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("Client connected with address : ", connection.RemoteAddr())

		buf := make([]byte, 512)

		connection.Read(buf)
		if err != nil {
			fmt.Println("Error while reading : ", err.Error())
		}

		fmt.Println("Received data from client : ", string(buf[:]))
	}
}

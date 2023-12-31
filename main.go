package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	SERVER_HOST     = "localhost"
	SERVER_PORT     = "4949"
	SERVER_PROTOCOL = "tcp"
	BUFFER_SIZE     = 1024
)

var serverHost string
var serverPort string
var serverProtocol string

func init() {
	flag.StringVar(&serverHost, "host", SERVER_HOST, "Server IP address")
	flag.StringVar(&serverPort, "port", SERVER_PORT, "Port for the Server")
	flag.StringVar(&serverProtocol, "protocol", SERVER_PROTOCOL, "Protocol for the Server")
}

func main() {

	flag.Parse()

	startServer()
}

func startServer() {
	address := fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT)

	server, err := net.Listen(SERVER_PROTOCOL, address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Server Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Printf("Server Waiting for Clients...\n\n")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		handleClient(connection)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected with address : ", conn.RemoteAddr())

	buf := make([]byte, BUFFER_SIZE)

	for {
		err := readData(conn, buf)
		if err != nil {
			break
		}
	}
}

func readData(conn net.Conn, buf []byte) error {

	dataLen, err := conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Client ", conn.RemoteAddr(), " disconnected : ", err.Error())
			return io.EOF
		}
	}

	data := string(buf[:dataLen])

	fmt.Println("Received from client : ", conn.RemoteAddr(), " and data : ", data)

	_, err = conn.Write([]byte("Got your message : " + data))
	if err != nil {
		fmt.Println("Error Writing:", err.Error())
		return err
	}

	return nil
}

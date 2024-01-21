package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func discoverTCPServerAddress() string {
	return "10.100.23.129"
}

func connectToTCPServer(serverAddr string, useFixedSize bool) {
	var conn net.Conn

	if useFixedSize {
		conn, _ = net.Dial("tcp", serverAddr+":34933")
	} else {
		conn, _ = net.Dial("tcp", serverAddr+":33546")
	}
	defer conn.Close()

	// Read welcome message
	buffer := make([]byte, 1024)
	conn.Read(buffer)

	fmt.Println("Server says:", string(buffer))

	// Send a message
	message := "Hello, TCP Server!\x00" // Null-terminated message
	conn.Write([]byte(message))

	// Read echo
	conn.Read(buffer)

	fmt.Println("Echo from server:", string(buffer))
}

func instructServerToConnectBack(serverAddr string) {
	conn, err := net.Dial("tcp", serverAddr+":33546") // Assuming using delimited messages
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := getLocalAddress()                                // Implement this function to get local IP
	message := fmt.Sprintf("Connect to: %s:20005\x00", localAddr) // Replace 20000 with your listening port
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func getLocalAddress() string {
	// Implement logic to get local IP address
	// This can be complex and environment-dependent
	return "10.100.23.20" // Placeholder for the actual implementation
}

func startReverseConnectionServer() {
	listener, _ := net.Listen("tcp", "10.100.23.20:20005") // Use the same IP and port as in instructServerToConnectBack

	defer listener.Close()

	for {
		conn, _ := listener.Accept()

		go handleConnection(conn) // Implement this function to handle the connection
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, _ := reader.ReadString('\x00') // Read until the null character

		// Process the message (in this case, just log it)
		fmt.Printf("Received message: %s\n", message)

		// Echo the message back to the sender
		conn.Write([]byte(message))

	}
}

func main() {
	serverAddr := discoverTCPServerAddress()

	go connectToTCPServer(serverAddr, false) // false for delimited messages
	go instructServerToConnectBack(serverAddr)
	startReverseConnectionServer()
}

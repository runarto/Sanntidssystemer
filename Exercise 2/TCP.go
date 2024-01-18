package main

import (
	"net"
	"fmt"
	"log"
	"bufio"
	"io"
)
func discoverTCPServerAddress() (string, error) {
    return "10.100.23.129", nil
}

func connectToTCPServer(serverAddr string, useFixedSize bool) {
    var conn net.Conn
    var err error
    if useFixedSize {
        conn, err = net.Dial("tcp", serverAddr+":34933")
    } else {
        conn, err = net.Dial("tcp", serverAddr+":33546")
    }
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // Read welcome message
    buffer := make([]byte, 1024)
    _, err = conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(buffer))

    // Send a message
    message := "Hello, TCP Server!\x00" // Null-terminated message
    _, err = conn.Write([]byte(message))
    if err != nil {
        log.Fatal(err)
    }

    // Read echo
    _, err = conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Echo from server:", string(buffer))
}

func instructServerToConnectBack(serverAddr string) {
    conn, err := net.Dial("tcp", serverAddr+":33546") // Assuming using delimited messages
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := getLocalAddress() // Implement this function to get local IP
    message := fmt.Sprintf("Connect to: %s:20000\x00", localAddr) // Replace 20000 with your listening port
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
    listener, err := net.Listen("tcp", "10.100.23.20:20000") // Use the same IP and port as in instructServerToConnectBack
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go handleConnection(conn) // Implement this function to handle the connection
    }
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\x00') // Read until the null character
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %v\n", err)
			}
			break
		}

		// Process the message (in this case, just log it)
		fmt.Printf("Received message: %s\n", message)

		// Echo the message back to the sender
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Printf("Error writing to connection: %v\n", err)
			break
		}
	}
}


func main() {
    serverAddr, err := discoverTCPServerAddress()
	
    if err != nil {
        log.Fatal(err)
    }

    go connectToTCPServer(serverAddr, false) // false for delimited messages
    go instructServerToConnectBack(serverAddr)
    startReverseConnectionServer()
}

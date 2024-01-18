package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

var serverIP = "10.100.23.129"
var localIP = "10.100.23.20"
var portFixedSizeMessage = 34933
var portZeroTerminatedMessage = 33546

func receiveMessage(conn *net.TCPConn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error here")
			fmt.Println("Error reading from TCP:", err)
			continue
		}
		message := string(buffer[:n])
		fmt.Println(message)
	}
}

func sendMSG(conn *net.TCPConn) {
	fmt.Println("Connection established")
	defer conn.Close()

	data := []byte("Connect to: 10.100.23.20:33546\x00")
	_, err := conn.Write(data)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Oh no, error")
	} else {
		fmt.Println("Message sent")
	}

	message := []byte("Hello world")

	for {
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Oh no, error")
		} else {
			fmt.Println("Message sent")
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	sendingAddress := &net.TCPAddr{
		Port: portZeroTerminatedMessage,
		IP:   net.ParseIP(serverIP),
	}

	connSendTCP, err := net.DialTCP("tcp", nil, sendingAddress)
	if err != nil {
		fmt.Println("Error creating TCP connection:", err)
		return
	}

	go receiveMessage(connSendTCP)
	go sendMSG((connSendTCP))

	time.Sleep(10 * time.Minute)
}

package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

var serverIP = "10.100.23.129"
var portSendMSG = 20010

func ListenForAddress(conn *net.UDPConn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		message := string(buffer[:n])
		fmt.Println(message)
	}
}

func sendMSG(conn *net.UDPConn) {
	fmt.Println("Connection established")
	defer conn.Close()

	data := []byte("Hello world!")
	fmt.Println("Data acquired")

	for {
		fmt.Println("I am here")
		_, err := conn.Write(data)
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

	sendingAddress := &net.UDPAddr{
		Port: portSendMSG,
		IP:   net.ParseIP("255.255.255.255"),
	}

	receivingAddr := &net.UDPAddr{
		Port: portSendMSG,
		IP:   net.ParseIP("0.0.0.0"),
	}

	connSend, err := net.DialUDP("udp", nil, sendingAddress)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	connReceive, _ := net.ListenUDP("udp", receivingAddr)

	go sendMSG(connSend)
	go ListenForAddress(connReceive)

	// Add a delay to keep the program running for a while
	time.Sleep(10 * time.Minute)
}

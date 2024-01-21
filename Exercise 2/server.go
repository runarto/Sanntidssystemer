package main

import (
    "fmt"
    "net"
)

func main() {
    addr := net.UDPAddr{
        Port: 30000,
        IP:   net.ParseIP("0.0.0.0"),
    }

    conn, err := net.ListenUDP("udp", &addr)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    buffer := make([]byte, 1024)

    for {
        _, remoteAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Error reading from UDP:", err.Error())
            continue
        }
        fmt.Println("Received message from IP:", remoteAddr.IP.String())
    }
}

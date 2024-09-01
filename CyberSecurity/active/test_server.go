package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

// Start a UDP server
func startUDPServer(port string) {
	addr := net.UDPAddr{
		Port: parsePort(port),
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("UDP server listening on port", port)

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		fmt.Printf("Received UDP message from %s: %s\n", remoteAddr, string(buf[:n]))
		_, err = conn.WriteToUDP([]byte("Hello from UDP server"), remoteAddr)
		if err != nil {
			fmt.Println("Error writing to UDP:", err)
		}
	}
}

// Start a TCP server
func startTCPServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("TCP server listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Accepted new TCP connection from", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Connection closed by client:", conn.RemoteAddr().String())
			} else {
				fmt.Println("Error reading from TCP connection:", err)
			}
			return
		}
		fmt.Printf("Received TCP message: %s\n", string(buf[:n]))
		_, err = conn.Write([]byte("Hello from TCP server"))
		if err != nil {
			fmt.Println("Error writing to TCP connection:", err)
			return
		}
	}
}

func parsePort(port string) int {
	p, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Invalid port:", port)
		os.Exit(1)
	}
	return p
}

func main() {
	udpPort := "8081"
	tcpPort := "8080"

	go startUDPServer(udpPort)
	go startTCPServer(tcpPort)

	select {} // Block forever
}

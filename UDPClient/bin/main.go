package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	// define server address
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	//connect to server
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	//send message to server
	message := "Hello BlackHatwithGo UDP Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	//create timeout so that Go doesn't sit there infinetly waiting for a response
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	//read response back from server and print it to screen
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading response from server", err)
		os.Exit(1)
	}

	fmt.Println("Server Response is:", string(buffer[:n]))

}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	server := "google.com:80"

	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	request := "GET / HTTP/1.1\r\n"
	request += "Host: google.com\r\n"
	request += "Connection: close\r\n"
	request += "\r\n"

	_, err = conn.Write([]byte(request))
	if err != nil {
		log.Fatalf("Failed to send request %v", err)
	}

	response := make([]byte, 4096)
	n, err := conn.Read(response)
	if err != nil && err != io.EOF {
		log.Fatalf("failed to read response %v", err)
	}

	fmt.Printf("Response from Google.com\n\n%s\n", string(response[:n]))
}

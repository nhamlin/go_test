package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// Set up a TCP listening server
func server() {
	ln, err := net.Listen("tcp", ":9999") // Listen on port 9999 for TCP traffic
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept() // Accept a connection
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(c) // Handle the connection
	}
}

// What do we do with the connection?
func handleServerConnection(c net.Conn) {
	var msg string                        // Receive the message
	err := gob.NewDecoder(c).Decode(&msg) // Decode the message
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg) // Print out the message that was received
	}
	c.Close()
}

func client(msg string) {
	c, err := net.Dial("tcp", "127.0.0.1:9999") // Connect to a server, localhost port 9999
	if err != nil {
		fmt.Println(err)
		return
	}
	// msg := "Hello, World"
	fmt.Println("Sending", msg)         // Print the message that is being sent
	err = gob.NewEncoder(c).Encode(msg) // Encode the message
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func main() {
	go server() // Create the server
	var input string
	fmt.Scanln(&input)
	go client(input) // Create the client with the message that was entered on the scan line
	fmt.Scanln(&input)
}

package main

import (
	"fmt"
	"log"
	"net"
)

func startServer(ip, port string) error {
	addr := fmt.Sprintf("%s:%s", ip, port)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()

	log.Println("Listening on", addr)

	err = acceptLoop(l)

	return err
}

func acceptLoop(l net.Listener) error {
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		// Handle connections in a new goroutine.
		go handleConnection(conn)
	}
}

// Handles incoming requests.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		// Make a buffer to hold incoming data.
		buf := make([]byte, 1)
		// Read the incoming connection into the buffer.
		bufLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		if bufLen == 1 {
			switch buf[0] {
			case 1:
				runCmd("play-pause")
			case 2:
				runCmd("next")
			case 3:
				runCmd("previous")
			}
		}
	}
}

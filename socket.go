package main

import (
	"fmt"
	"log"
	"net"
	"time"
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

var previous = ""

// Handles incoming requests.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	ackReceived := false
	death := false

	go func() {
		for range time.Tick(time.Second * 2) {
			if death {
				return
			}
			title := runCmd("metadata", "title")
			artist := runCmd("metadata", "artist")

			titleString := "Unknown"
			if title != nil {
				titleString = string(title)
			}

			artistString := "Unknown"
			if title != nil {
				artistString = string(artist)
			}

			finalString := fmt.Sprintf("%s by %s", titleString, artistString)
			log.Println(finalString)

			if previous == finalString {
				if ackReceived {
					continue
				}
			}

			finalUni := WriteUTF16String(finalString)

			finalUni = append([]byte{4}, finalUni...)

			_, err := conn.Write(finalUni)
			if err != nil {
				fmt.Println("Error writing:", err.Error())
				death = true
				return
			}

			ackReceived = false
			previous = finalString
		}
	}()

	for {
		if death {
			break
		}
		// Make a buffer to hold incoming data.
		buf := make([]byte, 1)
		// Read the incoming connection into the buffer.
		bufLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			death = true
			break
		}

		//handle commands
		if bufLen == 1 {
			switch buf[0] {
			case 1:
				runCmd("play-pause")
			case 2:
				runCmd("next")
			case 3:
				runCmd("previous")
			case 4:
				ackReceived = true
			}
		}
	}
}

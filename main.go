package main

import (
	"flag"
	"log"
)

func parsePort() string {
	port := flag.String("port", "8080", "--port 8080")

	flag.Parse()

	return *port
}

func main() {
	//parse port from cmd
	ip := parsePort()

	err := startServer("localhost", ip)
	if err != nil {
		panic(err)
	}

	log.Println("Shutting Down...")
}

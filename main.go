package main

import (
	"log"
	"net"
	"os"
)

var port string

func main() {
	s := newServer()

	go s.run()

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())

	}

	defer listener.Close()
	log.Printf("started server on:%s", port)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}

}

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	port := flag.Int("port", 8000, "Port to listen on")
	flag.Parse()
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("could not listen to tcp connections on port %d: %v", *port, err)
	} else {
		log.Printf("listening to tcp connections on localhost:%d", *port)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("could not accept connection: %v", err)
		}
		go func() {
			io.Copy(os.Stdout, conn)
		}()
	}
}

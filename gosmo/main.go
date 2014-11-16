package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"
)

type Data struct {
	D string
}

var (
	port string
)

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	d := &Data{}
	if err := dec.Decode(d); err != nil {
		log.Println("Decode error: ", err)
		return
	}
	log.Printf("Received : %+v\n", d)
}

func init() {
	flag.StringVar(&port, "p", "8080", "Port to listen to.")
	flag.Parse()
}

func main() {
	log.Println("Server started.")
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

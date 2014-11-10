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
		log.Fatal("Decode error: ", err)
	}
	log.Printf("Received : %+v\n", d)
}

func init() {
	flag.StringVar(&port, "p", "8080", "Port to listen to.")
	flag.Parse()
}

func main() {
	fmt.Println("start")
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
		// a goroutine handles conn so that the loop can accept other connections
		go handleConnection(conn)
	}
}

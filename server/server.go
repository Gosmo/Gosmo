package server

import (
	"encoding/gob"
	"log"
	"net"
)

type Data struct {
	D string
}

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	d := &Data{}
	if err := dec.Decode(d); err != nil {
		log.Println("Decode error: ", err)
		return
	}
	log.Printf("Received : %+v\n", d)
}

func Run(port string) {
	log.Println("Server started.")
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

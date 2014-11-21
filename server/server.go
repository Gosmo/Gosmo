package server

import (
	"encoding/gob"
	"log"
	"net"
)

type header struct {
	MachineID     string
	ContentLength int
}

type Data struct {
	Header header
	Body   string
}

var (
	machines = make(map[string][]string)
)

func receiveIntoMachines(id, data string) {
	if ds, ok := machines[id]; ok {
		ds = append(ds, data)
		machines[id] = ds
	} else {
		machines[id] = ds
		log.Printf("Added machine %s\n", id)
	}
	log.Printf("Received:\nMachine: %s\nData: %s\n", id, data)
}

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	d := &Data{}
	if err := dec.Decode(d); err != nil {
		log.Println("Decode error: ", err)
		return
	}
	id := d.Header.MachineID
	// For later use; only read the amount provided in ContentLength
	cntlen := d.Header.ContentLength
	if id == "" {
		log.Println("No machineID provided.")
		return
	}
	// Just a placeholder value.
	if cntlen > 1024 {
		log.Println("Too much content.")
		return
	}
	receiveIntoMachines(id, d.Body)
}

func Run(port string) {
	log.Println("Server started on port:", port)
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

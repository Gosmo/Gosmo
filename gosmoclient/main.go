package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

type Data struct {
	D string
}

// Sort of placeholder for properly doing this.
// Also, requires the full path of the file.
func runScript(file string) string {
	out, err := exec.Command("sh", file).Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func init() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Need a script to run.")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("start client")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := gob.NewEncoder(conn)
	rv := runScript(os.Args[1])
	err = encoder.Encode(Data{rv})
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
	fmt.Println("done")
}

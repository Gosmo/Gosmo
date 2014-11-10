package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

type Data struct {
	D string
}

var (
	port   string
	script string
)

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
	flag.StringVar(&port, "p", "8080", "Port to connect to.")
	flag.StringVar(&script, "s", "", "Script to run.")
	flag.Parse()

	if script == "" {
		fmt.Fprintln(os.Stderr, "Need a script to run.")
		os.Exit(1)
	}
}

func main() {
	// Localhost as a placeholder and for testing, to be a configurable option.
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := gob.NewEncoder(conn)
	rv := runScript(script)
	if err = encoder.Encode(Data{rv}); err != nil {
		log.Fatal(err)
	}
	conn.Close()
}

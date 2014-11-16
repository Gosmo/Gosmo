package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
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
	interp := inferLang(file)
	if interp == "" {
		fmt.Fprintln(os.Stderr, "Interpreter not inferred.")
		os.Exit(1)
	}
	out, err := exec.Command(interp, file).Output()
	if err != nil {
		return "exec error with: " + file
	}
	return string(out)
}

// Use the info of the file as well as this.
func inferLang(file string) string {
	switch {
	case strings.HasSuffix(file, ".sh"):
		return "sh"
	case strings.HasSuffix(file, ".lua"):
		return "lua"
	default:
		return ""
	}
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

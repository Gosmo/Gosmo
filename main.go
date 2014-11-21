package main

import (
	"fmt"
	"os"

	"github.com/Gosmo/client"
	"github.com/Gosmo/server"
)

const (
	serverMode = iota
	clientMode
)

var (
	gosmoMode  int
	host, port string
	scriptDir  string
	machineID  string
)

func letArgs(args []string, n int) string {
	if len(args) <= n {
		fmt.Fprintln(os.Stderr, "Insufficient arguments.")
		fmt.Println(args)
		usage()
	}
	n = n + 1
	return args[n]
}

func usage() {
	fmt.Println(`Usage:
    server -p[ort]
    client -id -p[ort] -host -s[cripts]
    -h[elp]`)
	os.Exit(0)
}

func init() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
	}

	switch args[0] {
	case "server":
		gosmoMode = serverMode
	case "client":
		gosmoMode = clientMode
	default:
		usage()
	}

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "-h", "-help", "--help":
			usage()
		case "-p", "-port", "--port":
			port = letArgs(args, i)
		case "-host", "--host":
			host = letArgs(args, i)
		case "-s", "-scripts", "--scripts":
			scriptDir = letArgs(args, i)
		case "-id", "--id":
			machineID = letArgs(args, i)
		}
	}

	switch "" {
	case port:
		port = "28080"
	case host:
		host = "localhost"
	case scriptDir:
		usage()
	}
}

func main() {
	switch gosmoMode {
	case serverMode:
		server.Run(port)
	case clientMode:
		client.Run(machineID, host, port, scriptDir)
	}
}

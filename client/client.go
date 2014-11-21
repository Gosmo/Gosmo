// TODO: take a dir as an arg, the dir contains the scripts and blobs to be run.
package client

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

type header struct {
	MachineID     string
	ContentLength int
}

type Data struct {
	Header header
	Body   string
}

// Sort of placeholder for properly doing this.
// Also, requires the full path of the file.
func runScript(file string) string {
	interp := inferLang(file)
	if interp == "" {
		fmt.Fprintln(os.Stderr, "Interpreter not inferred.")
		return "Interpreter not inferred with: " + file
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

func getScripts(scriptDir string) (files []string) {
	dir, err := os.Open(scriptDir)
	if err != nil {
		log.Printf("Failed to open %s with error: %v\n", scriptDir, err)
		return
	}
	defer dir.Close()

	files, err = dir.Readdirnames(0)
	if err != nil {
		log.Printf("Failed to read %s with error: %v\n", scriptDir, err)
	}
	return
}

func Run(machineID, host, port, scriptDir string) {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatal("Connection error with:", err)
	}
	defer conn.Close()

	encoder := gob.NewEncoder(conn)

	if !strings.HasSuffix(scriptDir, "/") {
		scriptDir += "/"
	}

	var allData string
	for _, script := range getScripts(scriptDir) {
		rv := runScript(scriptDir + script)
		allData += "\n" + rv
	}
	d := Data{
		Header: header{machineID, len(allData)},
		Body:   allData,
	}
	if err = encoder.Encode(d); err != nil {
		log.Println(err)
	}
}

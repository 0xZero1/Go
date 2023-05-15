package main

import (
	"net"
	"log"
	"io"
)

func echo(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("[-] Unable to read/write data")
	}
}

func main() {
	var port = ":20080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("[-] Unable to bind to port...")
	}
	log.Println("[+] Listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("[+] Received connection")
		if err != nil {
			log.Fatalln("[-] Unable to accept connection...")
		}
		go echo(conn)
	}
}

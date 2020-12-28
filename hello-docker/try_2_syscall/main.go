package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(syscall.SYS_GETPID, 0, 0, 0)
	fmt.Println("process id: ", pid)
	cmd := exec.Command("/bin/bash", "-c", "echo $PPID")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("get err when exec cmd", err)
		return
	}
	fmt.Println("get ppid is->", out.String())
}

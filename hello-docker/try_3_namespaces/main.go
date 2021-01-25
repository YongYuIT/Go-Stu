package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for true {
		str, _, err := in.ReadLine()
		if err != nil {
			fmt.Println("input err-->", err)
		}
		if strings.EqualFold(string(str), "Mount") {
			tryMount()
		}
		if strings.EqualFold(string(str), "UTS") {
			tryUTS()
		}
		if strings.EqualFold(string(str), "IPC") {
			tryIPC()
		}
		if strings.EqualFold(string(str), "PID") {
			tryPID()
		}
		if strings.EqualFold(string(str), "Network") {
			tryNetwork()
		}
		if strings.EqualFold(string(str), "User") {
			tryUser()
		}
	}
}

func tryUser() {

}

func tryNetwork() {

}

func tryPID() {

}

func tryIPC() {

}

/*
$ sudo su
$ go run main.go
uts
# pstree -pl
go(2867)─┬─main(2969)─┬─sh(2990)
# echo $$
2990
# readlink /proc/2867/ns/uts
uts:[4026531838]
# readlink /proc/2969/ns/uts
uts:[4026531838]
# readlink /proc/2990/ns/uts
uts:[4026532631]
# hostname
ubuntu
# hostname -b fucku
# hostname
fucku
$ hostname
ubuntu
*/
func tryUTS() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("start cmd err-->", err)
	}
}

func tryMount() {

}

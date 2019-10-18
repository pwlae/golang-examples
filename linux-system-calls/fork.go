package main

import (
	"syscall"
	"os"
)

func main() {
	println("PID: ", syscall.Getpid())
	println("PPID: ", os.Getppid())
}

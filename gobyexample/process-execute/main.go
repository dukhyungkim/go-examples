package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	if err := syscall.Exec(binary, args, env); err != nil {
		panic(err)
	}
}

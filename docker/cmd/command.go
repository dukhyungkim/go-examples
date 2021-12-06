package main

import (
	"bytes"
	"log"
	"os/exec"
)

func RunCmd(cmdStr string) string {
	cmd := exec.Command("/bin/bash", "-c", cmdStr)

	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	if err := cmd.Run(); err != nil {
		log.Printf("failed to run command: %v\n", err)
		log.Fatalln(stdErr.String())
	}

	return stdOut.String()
}

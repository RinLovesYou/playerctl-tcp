package main

import (
	"log"
	"os/exec"
)

func runCmd(commands ...string) []byte {
	cmd := exec.Command("playerctl", commands...)
	stdout, err := cmd.Output()

	if err != nil {
		log.Printf("Error executing playerctl: %s", err)
		return nil
	}

	return stdout
}

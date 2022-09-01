package main

import (
	"log"
	"os/exec"
)

func runCmd(command string) {
	cmd := exec.Command("playerctl", command)
	stdout, err := cmd.Output()

	if err != nil {
		log.Printf("Error executing playerctl: %s", err)
		return
	}

	log.Println(stdout)
}

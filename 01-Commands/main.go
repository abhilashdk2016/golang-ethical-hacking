package main

import (
	"os"
	"os/exec"
	"log"
)

func executeCommand(command string, args_array []string) (err error) {
	args := args_array
	cmd_obj := exec.Command(command, args...)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr

	err = cmd_obj.Run()
	if err != nil { 
		return err
	}
	return nil
}

func main() {
	command := "ls"
	err := executeCommand(command, []string{"-l"})
	if err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"flag"
	"os"
	"os/exec"
)

func executeCommand(command string, args_array []string) (err error) {
	args := args_array
	cmd_obj := exec.Command(command, args...)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr
	cmd_obj.Stdin = os.Stdin

	err = cmd_obj.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	iface := flag.String("iface", "eth0", "Interface for which you want to change the MAC")
	newMac := flag.String("newMac", "", "IProvide New Mac Address")

	flag.Parse()

	executeCommand("sudo", []string{"ifconfig", *iface, "down"})
	executeCommand("sudo", []string{"ifconfig", *iface, "hw", "ether", *newMac})
	executeCommand("sudo", []string{"ifconfig", *iface, "up"})
}

// Reset MAC Address
// macchanger -p eth0

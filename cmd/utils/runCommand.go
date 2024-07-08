package utils

import (
	"os"
	"os/exec"
)

// runCommand is a utility function to execute a system command
// It takes the command name and its arguments, runs it, and directs the output to stdout and stderr
func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

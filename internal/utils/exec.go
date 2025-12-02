package utils

import (
	"os/exec"
)

// Exec runs a command and captures its standard output.
// It returns the output as a string and an error if the command fails.
func Exec(c string, args ...string) (string, error) {
	cmd := exec.Command(c, args...)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(stdout), nil
}

// ExecV runs a command for its side effects, without capturing its output.
// The command's output is printed directly to the terminal.
// It returns an error if the command fails.
func ExecV(c string, args ...string) error {
	cmd := exec.Command(c, args...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

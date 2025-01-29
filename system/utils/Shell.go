package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// Cexec executes a command in the default shell and captures the standard output (stdout) and errors (stderr).
// It identifies the shell being used by checking the SHELL environment variable.
func Cexec(command string) (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "sh" // Default to "sh" if SHELL is not set
	}

	cmd := exec.Command(shell, "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return out.String(), fmt.Errorf("error when executing the command: %s\nstderr: %s", err.Error(), stderr.String())
	}

	if cmd.ProcessState.ExitCode() != 0 {
		return out.String(), fmt.Errorf("stderr: %s", stderr.String())
	}

	return out.String(), nil
}

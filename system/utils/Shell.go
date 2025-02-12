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
	// Check if the shell is executable
	if path, err := exec.LookPath(shell); err != nil {
		return "", fmt.Errorf("shell not found: %s", shell)
	} else {
		shell = path
	}
	cmd := exec.Command(shell, "-c", command)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return out.String(), fmt.Errorf("error when executing the command: %s\nstderr: %s", err.Error(), stderr.String())
	}

	if exitCode := cmd.ProcessState.ExitCode(); exitCode != 0 {
		return out.String(), fmt.Errorf("stderr: %s", stderr.String())
	}

	return out.String(), nil
}

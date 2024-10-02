package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Cexec(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)

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

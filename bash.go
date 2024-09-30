package lcme

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Bash(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return out.String(), fmt.Errorf("error executing command: %s\n%s", err, stderr.String())
	}

	if cmd.ProcessState.ExitCode() != 0 {
		return out.String(), fmt.Errorf("command exited with non-zero status: %d", cmd.ProcessState.ExitCode())
	}

	return out.String(), nil
}

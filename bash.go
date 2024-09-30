package lcme

import (
	"bytes"
	"os/exec"
)

func Bash(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

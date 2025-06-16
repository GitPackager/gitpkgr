package utils

import (
	"fmt"
	"os/exec"
)

func ExecuteCommand(args []string) (out string, e error) {
	cmd := exec.Command(args[0], args[1:]...)
	output, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Printf("ExecuteCommand: error: %v\n", e)
		return "", e
	}
	return string(output), nil
}

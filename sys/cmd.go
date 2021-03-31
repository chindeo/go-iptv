package sys

import (
	"os"
	"os/exec"
	"strings"
)

func CmdRun(name string, args []string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getPid(name string) string {
	output, _ := exec.Command("pgrep", "-f", name).Output()
	return strings.TrimSpace(string(output))
}

func IsRunning(name string) bool {
	return getPid(name) != ""
}

func Stop(name string) bool {
	cmd := exec.Command("kill", "-TERM", getPid(name))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

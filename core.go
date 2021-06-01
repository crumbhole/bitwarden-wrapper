package bitwardenwrapper

import (
	"os/exec"
)

func runBw(flags ...string) (string, error) {
	cmd := exec.Command("bw", flags...)
	out, err := cmd.Output()
	return string(out), err
}

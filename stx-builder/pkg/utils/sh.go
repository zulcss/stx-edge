package utils

import (
	"os"
	"os/exec"
)

func SH(c string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Env = os.Environ()
	o, err := cmd.CombinedOutput()
	return string(o), err
}

func CheckUser() bool {
	if os.Geteuid() != 0 {
		return false
	}
	return true
}
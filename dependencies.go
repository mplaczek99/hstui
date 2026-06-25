package main

import (
	"fmt"
	"os/exec"
)

func CheckDependencies() error {
	for _, bin := range []string{"hyprsunset", "uwsm", "notify-send"} {
		if _, err := exec.LookPath(bin); err != nil {
			return fmt.Errorf("%s not found in PATH", bin)
		}
	}

	return nil
}

func Notify(body string) error {
	return exec.Command(
		"notify-send",
		"-a", "hstui",
		"-u", "critical",
		"hstui",
		body,
	).Run()
}

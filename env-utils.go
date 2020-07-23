package main

import (
	"os"
	"os/exec"
	"strings"
)

func hasDirenv() bool {
	_, err := exec.LookPath("direnv")
	return err == nil
}

func getEnv() []string {
	if hasDirenv() {
		command := []string{"direnv", "exec", ".", "env"}
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Env = os.Environ()
		output, err := cmd.Output()
		if err != nil {
			fatalError(1, "%s", err)
		}
		env := strings.Split(string(output), "\n")
		return env
	} else {
		return os.Environ()
	}
}

func getEnvPlus(key string) string {
	for _, v := range getEnv() {
		parts := strings.SplitN(v, "=", 2)
		if parts[0] == key {
			return parts[1]
		}
	}
	return os.Getenv(key)
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// AWS_VAULT_PROFILE

func buildMainCommand(profile string) []string {
	mainCmd := os.Args[1:]
	if profile != "" {
		wrapper := []string{"aws-vault", "exec", profile, "--"}
		mainCmd = append(wrapper, mainCmd...)
	}
	return mainCmd
}

func fatalError(code int, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "[aws-vault-exec ERROR] %s\n", message)
	os.Exit(code)
}

func execCommand(args []string) {
	executable, err := exec.LookPath(args[0])
	if err != nil {
		fatalError(1, "%s: trying to LookPath %s", err, args[0])
	}

	err = syscall.Exec(executable, args, os.Environ())
	if err != nil {
		fatalError(1, "%s: trying to exec %s", err, args)
	}
}

func main() {
	profile := os.Getenv("AWS_VAULT_PROFILE")
	mainCmd := buildMainCommand(profile)
	if len(mainCmd) == 0 {
		fatalError(1, "no command specified")
	}
	execCommand(mainCmd)
}

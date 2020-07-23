package main

import (
	"os/exec"
	"syscall"
)

func execCommand(args []string) {
	executable, err := exec.LookPath(args[0])
	if err != nil {
		fatalError(1, "%s: trying to LookPath %s", err, args[0])
	}

	err = syscall.Exec(executable, args, getEnv())
	if err != nil {
		fatalError(1, "%s: trying to exec %s", err, args)
	}
}

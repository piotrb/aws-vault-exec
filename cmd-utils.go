package main

import "os"

func buildMainCommand(profile string) []string {
	mainCmd := os.Args[1:]
	if profile != "" {
		wrapper := []string{"aws-vault", "exec", profile, "--"}
		mainCmd = append(wrapper, mainCmd...)
	}
	return mainCmd
}

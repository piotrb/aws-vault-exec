package main

import (
	"fmt"
	"os"
)

// AWS_VAULT_PROFILE

func main() {
	if len(os.Args) == 0 {
		fatalError(1, "no command specified")
	}

	profile := getEnvPlus("AWS_VAULT_PROFILE")
	mainCmd := buildMainCommand(profile)
	fmt.Printf("%s\n", mainCmd)
	execCommand(mainCmd)
}

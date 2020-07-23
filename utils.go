package main

import (
	"fmt"
	"os"
)

func fatalError(code int, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "[aws-vault-exec ERROR] %s\n", message)
	os.Exit(code)
}

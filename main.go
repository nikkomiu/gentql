package main

import (
	"context"
	"os"

	"github.com/nikkomiu/gentql/cmd"
	"github.com/nikkomiu/gentql/pkg/errors"
)

func main() {
	ctx := context.Background()

	if err := cmd.Execute(ctx); err != nil {
		var exitCode int
		switch typedErr := err.(type) {
		case errors.ExitCodeError:
			exitCode = typedErr.ExitCode()

		default:
			exitCode = 1
		}

		os.Exit(exitCode)
	}

}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nikkomiu/gentql/cmd"
)

func main() {
	ctx := context.Background()

	if err := cmd.Execute(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute gentql: %s\n", err)
		os.Exit(1)
	}
}

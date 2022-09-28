package main

import (
	"fmt"
	"os"

	saga "github.com/einride/saga/cmd/saga/gen"
)

func main() {
	if err := saga.NewModuleCommand("saga", "Saga CLI").Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

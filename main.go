package main

import (
	"fmt"
	"log"

	"github.com/avrebarra/minimok/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		err = fmt.Errorf("unexpected error: %w", err)
		log.Panic(err)
		return
	}
}

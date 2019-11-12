package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)

	if err := run(); err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
}

// run ...
func run() error {
	err := FlagSet().Run(os.Args)
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"log"

	"github.com/muhammadalica4/devkit/internal/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}
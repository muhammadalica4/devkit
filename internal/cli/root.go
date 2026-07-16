package cli

import (
	"fmt"
	"os"
	
	"github.com/muhammadalica4/devkit/internal/commands/version"
	"github.com/muhammadalica4/devkit/internal/commands/uuid"
)

var registry = NewRegistry()

func Run() error {

	if len(os.Args) < 2 {
		printHelp()
		return nil
	}

	name := os.Args[1]

	return registry.Execute(name, os.Args[2:])
}

func init() {
    registry.Register(version.New())
	registry.Register(uuid.New())
}

func printHelp() {

	fmt.Println("Developer Toolbox")
	fmt.Println()

	fmt.Println("Available Commands")

	for _, cmd := range registry.List() {
		fmt.Printf("  %-15s %s\n",
			cmd.Name(),
			cmd.Description(),
		)
	}
}
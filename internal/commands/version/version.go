package version

import "fmt"

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c *Command) Name() string {
	return "version"
}

func (c *Command) Description() string {
	return "Print application version"
}

func (c *Command) Run(args []string) error {
	fmt.Println("devtool v0.1.0")
	return nil
}
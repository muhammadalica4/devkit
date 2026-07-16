package uuid

import (
	"crypto/rand"
	"fmt"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c *Command) Name() string {
	return "uuid"
}

func (c *Command) Description() string {
	return "Generate UUID v4"
}

func (c *Command) Run(args []string) error {

	u := make([]byte, 16)

	_, err := rand.Read(u)
	if err != nil {
		return err
	}

	u[6] = (u[6] & 0x0f) | 0x40
	u[8] = (u[8] & 0x3f) | 0x80

	fmt.Printf("%x-%x-%x-%x-%x\n",
	    u[0:4],
	    u[4:6],
	    u[6:8],
	    u[8:10],
	    u[10:16],
	)

	return nil
}
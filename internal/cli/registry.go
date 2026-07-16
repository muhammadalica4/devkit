package cli

import "fmt"

type Registry struct {
	commands map[string]Command
}

func NewRegistry() *Registry {
	return &Registry{
		commands: make(map[string]Command),
	}
}

func (r *Registry) Register(cmd Command) {
	r.commands[cmd.Name()] = cmd
}

func (r *Registry) Get(name string) (Command, bool) {
	cmd, ok := r.commands[name]
	return cmd, ok
}

func (r *Registry) List() []Command {
	out := make([]Command, 0, len(r.commands))

	for _, c := range r.commands {
		out = append(out, c)
	}

	return out
}

func (r *Registry) Execute(name string, args []string) error {
	cmd, ok := r.Get(name)

	if !ok {
		return fmt.Errorf("unknown command: %s", name)
	}

	return cmd.Run(args)
}
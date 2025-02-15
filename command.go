package main

import "errors"

type Command struct {
	Name string
	Args []string
}

type commands struct {
	registerCommands map[string]func(*state, Command) error
}

func (c *commands) Register(name string, f func(*state, Command) error) {
	c.registerCommands[name] = f
}

func (c *commands) Run(s *state, cmd Command) error {
	f, ok := c.registerCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

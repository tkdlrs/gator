package config

import (
	"errors"
	"fmt"
)

type state struct {
	cfgPointer *Config
}

type command struct {
	name string
	args []string
}

// handlers
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("handler expects an argument")
	}
	//
	err := s.cfgPointer.SetUser(cmd.args[1])
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}
	fmt.Println("user has been set")
	return nil
}

type commands struct {
	name map[string]func(*state, command) error
}

// Runs a given command with the provided state
func (c *command) run(s *state, cmd command) error {
	// splitArgs := strings.Split(cmd.args, " ")
	splitArgs := cmd.args
	if len(splitArgs) > 2 {
		return fmt.Errorf("ERROR: args less than two")
	}
	// need to run the given command
	// splitArgs[0](splitArgs[1])
	fmt.Println("need to run the given command")
	return nil

}

// Method registers a new handler function for a command name
func (c *command) register(name string, f func(*state, command) error) {
}

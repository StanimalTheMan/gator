package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/StanimalTheMan/gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	_, ok := c.commandsMap[cmd.name]
	if ok {
		c.commandsMap[cmd.name](s, cmd)
		return nil
	}

	return errors.New("command does not exist")
}

func handlerLogin(s *state, cmd command) error {
	fmt.Println("LOGIN", cmd.args)
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}
	fmt.Println("cmd args", cmd.args)
	s.config.SetUser(cmd.args[0])
	fmt.Println("username has been set")
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	st := state{config: &cfg}
	cmds := commands{
		commandsMap: make(map[string]func(*state, command) error),
	}

	cmds.commandsMap["login"] = handlerLogin

	userArgs := os.Args

	if len(userArgs) < 2 {
		fmt.Println("fewer than 2 args.  error...")
		os.Exit(1)
	}

	cmd := command{
		userArgs[1],
		userArgs[2:],
	}

	err = cmds.commandsMap[userArgs[1]](&st, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// err = cfg.SetUser("stanimal")
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}

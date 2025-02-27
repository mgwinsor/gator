package main

import (
	"log"
	"os"

	"github.com/mgwinsor/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	userCmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(programState, userCmd)
	if err != nil {
		log.Fatal(err)
	}
}

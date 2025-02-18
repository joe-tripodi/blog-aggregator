package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joe-tripodi/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	st := &state{
		cfg: &cfg,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("expecting 2 arguments")
		os.Exit(1)
	}
	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = commands.run(st, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

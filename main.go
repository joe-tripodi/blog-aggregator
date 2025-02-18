package main

import (
	"fmt"

	"github.com/joe-tripodi/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config:", err)
	}
	fmt.Println(cfg)
	err = cfg.SetUser("joe")
	if err != nil {
		fmt.Println("error writing to config:", err)
	}
	cfg, _ = config.Read()
	fmt.Println(cfg)
}

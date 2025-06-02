package main

import (
	"log"

	"github.com/tkdlrs/gator/internal/config"
)

func main() {
	// cfg, err := config.Read()
	// if err != nil {
	// 	log.Fatalf("error reading config: %v", err)
	// }
	// fmt.Printf("Read config: %+v\n", cfg)
	// //
	// err = cfg.SetUser("levi")
	// if err != nil {
	// 	log.Fatalf("Could not set current user: %v", err)
	// }
	// //
	// cfg, err = config.Read()
	// if err != nil {
	// 	log.Fatalf("error reading config: %v", err)
	// }
	// fmt.Printf("Read config again: %+v\n", cfg)

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	state := State{&cfg}

	theCommands := commands{
		"login": handlerLogin,
	}
}

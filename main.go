package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/tkdlrs/gator/internal/config"
	"github.com/tkdlrs/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	//
	dbURL := "postgres://postgres:postgres@localhost:5432/gator"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error rctabase URL: %v", err)
	}
	dbQueries := database.New(db)
	//
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	//
	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}
	//
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	//
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}
	//
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	//
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}

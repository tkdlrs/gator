package main

import (
	"context"
	"database/sql"
	"fmt"
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
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	//
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	//
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	//
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("feeds", handlerListFeeds)
	// Middleware ensures a user is logged in
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	//
	cmds.register("browse", middlewareLoggedIn(handlerListPosts))
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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// Get the ID of the current user.
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't find user: %w", err)
		}
		//
		return handler(s, cmd, user)
	}
}

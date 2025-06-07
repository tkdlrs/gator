package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tkdlrs/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	//
	userArgs := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}
	user, err := s.db.CreateUser(context.Background(), userArgs)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	//
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	//
	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	//
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	//
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	//
	fmt.Println("User switched successfully!")
	return nil
}

func printUser(user database.User) {
	fmt.Printf("	* ID:		%v\n", user.ID)
	fmt.Printf("	* Name:		%v\n", user.Name)
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to delete all users: %v", err)
	}
	//
	fmt.Println("Successfully removed all users")
	return nil
}

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
	/*
		ID
		CreatedAt
		UpdatedAt
		Name
	*/
	userArgs := database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name}
	_, err := s.db.CreateUser(context.Background(), userArgs)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	//
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	//
	fmt.Println("User successfully created")
	return nil
}

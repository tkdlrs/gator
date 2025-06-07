package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to delete all users: %v", err)
	}
	//
	fmt.Println("Successfully removed all users")
	return nil
}

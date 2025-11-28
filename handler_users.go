package main

import (
	"fmt"
	"context"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving users: %w", err)
	}

	for _, user := range users {
		if user == s.cfg.CurrentUserName {
			fmt.Printf(" * %v (current)\n", user)
		} else {
			fmt.Printf(" * %v\n", user)
		}
	}
	return nil
}
package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting users: %w", err)
	}
	err = s.db.DeleteFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting feeds: %w", err)
	}
	return nil
}

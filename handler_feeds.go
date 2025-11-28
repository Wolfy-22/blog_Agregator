package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	for _, feed := range feeds {
		username, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting username")
		}
		fmt.Printf(" Feed name: %v\n", feed.Name)
		fmt.Printf(" Feed URL: %v\n", feed.Url)
		fmt.Printf(" Username: %v\n\n", username)
	}
	return nil
}

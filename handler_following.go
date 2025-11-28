package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting followed feeds: %w", err)
	}

	for _, feed := range feedFollows {
		fmt.Printf(" * %v\n", feed.FeedName)
	}
	return nil
}

package main

import (
	"blog_Agregator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}
	userID, err := s.db.GetUserId(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user id: %w", err)
	}

	followFeed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating Feed Follow: %w", err)
	}

	fmt.Printf("Feed Name: %v\nUser Name: %v\n", followFeed.FeedName, followFeed.UserName)
	return nil
}

package main

import (
	"blog_Agregator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("add feed reuquires a name and a url")
	}

	currentUserId, err := s.db.GetUserId(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user data: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currentUserId,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	followFeed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentUserId,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating Feed Follow: %w", err)
	}

	fmt.Printf("%v\n", feed)
	fmt.Printf("%v\n", followFeed)
	return nil
}

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/StanimalTheMan/gator/internal/database"
)

func addFeed(s *state, cmd command) error {
	userName := s.cfg.CurrentUserName
	userInputs := os.Args
	fmt.Println(len(userInputs))
	if len(userInputs) < 4 {
		return errors.New("not enough arguments provided")
	}
	// if len(userInputs) != 2 {
	// 	return errors.New("only provide error")
	// }
	feedName := userInputs[2]
	feedUrl := userInputs[3]
	fmt.Println("feedName", feedName)
	fmt.Println("feedUrl", feedUrl)
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return errors.New("user not found")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed%w", err)
	}
	fmt.Printf("%+v", feed)
	return nil
}

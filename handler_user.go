package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/StanimalTheMan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	Context := context.Background()

	// check if user exists before logging in
	user, err := s.db.GetUser(Context, name)
	if err != nil {
		return fmt.Errorf("error occurred while fetching user")
	}
	if user.ID == uuid.Nil {
		log.Fatalf("user not found")
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	Context := context.Background()
	newUserID := uuid.New()
	name := cmd.Args[0]

	// create user
	user, err := s.db.CreateUser(Context, database.CreateUserParams{
		ID:        newUserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})

	if err != nil {
		log.Fatalf("failed to get user")
		return err
	}
	if user.ID == uuid.Nil {
		log.Fatalf("Name already exists")
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
		return err
	}
	fmt.Printf("user created %v", user)
	return nil
}

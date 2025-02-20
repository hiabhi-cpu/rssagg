package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hiabhi-cpu/rssagg/internal/database"
)

func repl() {

}

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("No arguments present")
	}
	ctx := context.Background()
	oldUser, err := s.db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("User doesnot exist")
	}
	err = s.con.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	s.con.Current_user_name = oldUser.UserName
	// fmt.Println("User switched successfully!")
	fmt.Printf("	-'%v'\n", s.con.Current_user_name)
	return nil
}

func handlerRegister(s *state, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("No arguments present")
	}
	ctx := context.Background()
	oldUser, err := s.db.GetUser(ctx, cmd.Args[0])
	if oldUser.UserName != "" {
		os.Exit(1)
	}
	newUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserName:  cmd.Args[0],
	})
	if err != nil {
		fmt.Errorf("Error in updating user", err)
	}
	s.con.SetUser(newUser.UserName)
	s.con.Current_user_name = newUser.UserName
	fmt.Println("New User is created")
	fmt.Println(newUser)
	return nil
}

func handlerReset(s *state, cmd Command) error {
	ctx := context.Background()
	s.db.DeleteAll(ctx)
	return nil
}

func handlerUsers(s *state, cmd Command) error {
	ctx := context.Background()
	users, err := s.db.ListUsers(ctx)
	if err != nil {
		return err
	}
	for _, u := range users {
		if u == s.con.Current_user_name {
			fmt.Printf("\t* %v (current)\n", u)
		} else {
			fmt.Printf("\t* %v\n", u)
		}

	}
	return nil
}

func handlerAgg(s *state, cmd Command) error {
	ctx := context.Background()
	// fmt.Println("hello")
	rss, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	for _, r := range rss.Channel.Item {
		fmt.Println(r.Title)
	}
	fmt.Println("Optimize for simplicity")
	return nil
}

func handlerAddFeed(s *state, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Not enough arguments present")
	}
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.con.Current_user_name)
	if err != nil {
		return fmt.Errorf("Error in getting current user", err)
	}
	newFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedName:  cmd.Args[0],
		FeedUrl:   cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		fmt.Errorf("Error in updating user", err)
	}
	// s.con.SetUser(newUser.UserName)
	// s.con.Current_user_name = newUser.UserName
	fmt.Println("New Feed is created")
	fmt.Println(newFeed)
	return nil
}

func handlerFeeds(s *state, cmd Command) error {
	ctx := context.Background()
	feeds, err := s.db.ListFeeds(ctx)
	if err != nil {
		return err
	}
	for _, f := range feeds {
		user, err := s.db.GetUserById(ctx, f.UserID)
		if err != nil {
			return err
		}
		fmt.Println(f.FeedName, "\t", f.FeedUrl, "\t", user.UserName)
	}
	return nil
}

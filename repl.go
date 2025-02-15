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

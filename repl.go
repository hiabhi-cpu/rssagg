package main

import "fmt"

func repl() {

}

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("No arguments present")
	}
	err := s.con.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	s.con.Current_user_name = cmd.Args[0]
	fmt.Println("User switched successfully!")
	fmt.Printf("	-'%v'\n", s.con.Current_user_name)
	return nil
}

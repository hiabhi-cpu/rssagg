package main

import "fmt"

func repl() {

}

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("No arguments present")
	}
	s.con.Current_user_name = cmd.Args[0]
	fmt.Println("User has been set")
	return nil
}

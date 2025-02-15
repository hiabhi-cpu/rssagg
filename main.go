package main

import (
	"log"
	"os"

	"github.com/hiabhi-cpu/rssagg/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cfg, err := config.ReadConfig()
	// fmt.Println(os.UserHomeDir())
	if err != nil {
		println("Error in reading the file", err)
		return
	}

	programState := &state{
		con: &cfg,
	}

	cmds := commands{
		registerCommands: make(map[string]func(*state, Command) error),
	}

	cmds.Register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	// fmt.Println(programState)
	err = cmds.Run(programState, Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(con)
	// err = con.SetUser("abhi")
	// if err != nil {
	// 	println("Error in setting user in file", err)
	// 	return
	// }
	// con, err = config.ReadConfig()
	// if err != nil {
	// 	println("Error in reading the file", err)
	// 	return
	// }
	// fmt.Println(con)
	// port := os.Getenv("PORT") //getting port from os
	// if port == "" {
	// 	log.Fatal("No port present")
	// }
	// // fmt.Println(port)
	// router := chi.NewRouter()
	// srv := &http.Server{ //http server
	// 	Handler: router,
	// 	Addr:    ":" + port,
	// }
	// log.Printf("Server stating on port %v", port)
	// err := srv.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

package main

import (
	"fmt"

	"github.com/hiabhi-cpu/rssagg/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	con, err := config.ReadConfig()
	if err != nil {
		println("Error in reading the file", err)
		return
	}
	fmt.Println(con)
	err = con.SetUser("abhi")
	if err != nil {
		println("Error in setting user in file", err)
		return
	}
	con, err = config.ReadConfig()
	if err != nil {
		println("Error in reading the file", err)
		return
	}
	fmt.Println(con)
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

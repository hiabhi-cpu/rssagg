package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func getReadData(file string) (string, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}

	b1 := make([]byte, len(dat))
	n1, err := f.Read(b1)
	if err != nil {
		return "", err
	}
	// fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	// fmt.Println(string(b1[:n1]))
	return string(b1[:n1]), nil
}

func ReadConfig() (Config, error) {
	// fmt.Println(os.UserHomeDir())
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	path := fmt.Sprint(home, "/.gatorconfig.json")
	// fmt.Println(path)
	jsonData, err := getReadData(path)
	if err != nil {
		return Config{}, fmt.Errorf("Error in reading the file", err)
	}
	// fmt.Println(jsonData)
	var con Config
	if err := json.Unmarshal([]byte(jsonData), &con); err != nil {
		return Config{}, fmt.Errorf("Error in reading json", err)
	}
	// fmt.Println(con)
	return con, nil
}

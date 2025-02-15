package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func (c *Config) SetUser(user_name string) error {
	c.Current_user_name = user_name
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error converting to json format")
	}
	// fmt.Println(string(data))
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := fmt.Sprint(home, "/.gatorconfig.json")
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("Error while writing to a file")
	}

	return nil
}

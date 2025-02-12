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
	fmt.Println(string(data))
	err = os.WriteFile("gatorconfig.json", data, 0644)
	if err != nil {
		return fmt.Errorf("Error while writing to a file")
	}
	return nil
}

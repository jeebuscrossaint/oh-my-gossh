package app

import (
	"fmt"
	"os"
	"path/filepath"
)

type Title struct {
	Name     string `json:"name"`
	Subtitle string `json:"subtitle"`
	Tab      string `json:"tab"`
}

type SSH struct {
	Status bool   `json:"status"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
}

type Color struct {
	Active   string `json:"active"`
	Inactive string `json:"inactive"`
}

type Config struct {
	Title  Title          `json:"title"`
	SSH    SSH            `json:"ssh"`
	Color  Color          `json:"color"`
	Extras map[string]any `json:"-"`
}

func Exists() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return false
	}

	configPath := filepath.Join(homeDir, ".config", "ohmygossh", "gossh.toml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// File does not exist, create it
		err = os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			fmt.Println("Error creating directories:", err)
			return false
		}

		file, err := os.Create(configPath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return false
		}
		defer file.Close()
		fmt.Println("Config file created:", configPath)
		return true
	}

	fmt.Println("Config file exists:", configPath)
	Parse()
	return true
}

func Parse() {

}

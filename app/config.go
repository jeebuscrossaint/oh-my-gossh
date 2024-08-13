package app

import (
	"fmt"
	"os"
	"path/filepath"
)

type title struct {
	name     string
	subtitle string
	tab      string
}

type ssh struct {
	status bool
	host   string
	port   int
}

type color struct {
	active   string
	inactive string
}

func Exists() bool {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "ohmygossh")
	matches, err := filepath.Glob(filepath.Join(configDir, "*.conf"))
	if err != nil {
		fmt.Println("Error:", err)
		return false
	} else if len(matches) == 1 {
		Parse(matches[0])
		return true
	}
	return false
}

func Parse(filePath string) {
	fmt.Printf("Parsing config file: %s\n", filePath)

}

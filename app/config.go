package app

import (
	"fmt"
	"os"
	"path/filepath"
)

func ConfigExists() bool {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "ohmygossh")
	matches, err := filepath.Glob(filepath.Join(configDir, "*.conf"))
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return len(matches) == 1
}

func Parse() {

}

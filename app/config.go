package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Title struct {
	Name     string
	Subtitle string
	Tab      string
}

type SSH struct {
	Status int
	Host   string
	Port   int
}

type Color struct {
	Active   string
	Inactive string
}

type Config struct {
	Title  Title
	SSH    SSH
	Color  Color
	Extras map[string]any `toml:"-"`
}

func Parse() Config {
	var config Config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configFile := filepath.Join(homeDir, ".config", "ohmygossh", "gossh.toml")
	_, err = toml.DecodeFile(configFile, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return config
}

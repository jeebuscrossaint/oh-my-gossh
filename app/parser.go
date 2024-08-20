package app

import (
	"fmt"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
)

var GlobalConfig Config

type Config struct {
	Title    Title
	SSH      SSH
	Color    Color
	Projects map[string]Project
}

type Title struct {
	Page     string
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
	Letter   string
}

type Project struct {
	File  string
	Name  string
	About string
}

func LoadConfig() (Config, error) {
	var config Config
	var filePath string
	if runtime.GOOS == "windows" {
		filePath = os.ExpandEnv("%USERPROFILE%\\.config\\ohmygossh\\gossh.toml")
	} else {
		filePath = os.ExpandEnv("$HOME/.config/ohmygossh/gossh.toml")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file: %v", err)
	}

	if _, err := toml.Decode(string(data), &config); err != nil {
		return Config{}, fmt.Errorf("error decoding TOML: %v", err)
	}

	return config, nil
}

func InitConfig() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}
	GlobalConfig = config
	return nil
}

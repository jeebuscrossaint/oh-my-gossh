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
	Style    Style
	Projects map[string]Project
}

type Style struct {
	MDPath string `toml:"md_style"` // path to markdown style
}

type Title struct {
	Page     string
	Name     string
	AsciiArt string `toml:"ascii_file"` // path to ascii art if wanted.
	Subtitle string
	Tab      string
	Pages    []string `toml:"pages"` // list of pages
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
    
	// Load ASCII art - now required
	asciiPath := os.ExpandEnv(config.Title.AsciiArt)
	asciiData, err := os.ReadFile(asciiPath)
	if err != nil {
	    return Config{}, fmt.Errorf("error reading ASCII art file: %v", err)
	}
	config.Title.AsciiArt = string(asciiData)
    
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

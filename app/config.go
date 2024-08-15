package app

import (
	"fmt"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title    Title
	SSH      SSH
	Color    Color
	Projects map[string]Project
}

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

type Project struct {
	File  string
	Name  string
	About string
}

func Parse() {
	var filePath string
	if runtime.GOOS == "windows" {
		filePath = os.ExpandEnv("%USERPROFILE%\\.config\\ohmygossh\\gossh.toml")
	} else {
		filePath = os.ExpandEnv("$HOME/.config/ohmygossh/gossh.toml")
	}

	fmt.Printf("Reading config from: %s\n", filePath)

	// Read the file contents
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Print the raw contents of the file
	//fmt.Println("Raw TOML content:")
	//fmt.Println(string(data))

	var config Config
	if _, err := toml.Decode(string(data), &config); err != nil {
		fmt.Printf("Error decoding TOML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Title: %+v\n", config.Title)
	fmt.Printf("SSH: %+v\n", config.SSH)
	fmt.Printf("Color: %+v\n", config.Color)

	if config.Projects == nil {
		fmt.Println("Projects is nil")
	} else {
		fmt.Printf("Projects: %+v\n", config.Projects)
		for key, project := range config.Projects {
			fmt.Printf("Project [%s]: %+v\n", key, project)
		}
	}

	// Print the entire config struct
	//fmt.Printf("Full config: %+v\n", config)
}

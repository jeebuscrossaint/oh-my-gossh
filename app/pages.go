package app

import (
	"oh-my-gossh/internal"
)

func Tui() {

	err := InitConfig()
	internal.Error(err, "TOML is broken, check your config file", true)

	tabs := []string{"main.go", "projects.cc", "about.rs", "contact.sh"}

	// todo initalize project list for projcets list

}

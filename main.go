package main

import (
	"fmt"
	"oh-my-gossh/app"
	"os"
)

func main() {
	if err := app.InitConfig(); err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("%s\n", app.GlobalConfig.Title.Name)

}

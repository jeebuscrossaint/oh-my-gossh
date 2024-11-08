package main

import (
	"oh-my-gossh/app"
	
	"os"
	"fmt"
)

func main() {
	if err := app.InitConfig(); err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

}

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

	if app.GlobalConfig.SSH.Status == 1 {
		app.SSHExec()
	} else if app.GlobalConfig.SSH.Status == 0 {
		app.Exec()
	} else {
		fmt.Println("Invalid SSH status")
		os.Exit(1)
	}

}

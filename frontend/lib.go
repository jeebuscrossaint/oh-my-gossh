package frontend

import (
	"fmt"
	"os"
)

func Error(e error, check string, fatal bool) {
	if e != nil {
		fmt.Printf("Error: %v: %v", check, e)
		if fatal {
			os.Exit(1)
		}
	}
}

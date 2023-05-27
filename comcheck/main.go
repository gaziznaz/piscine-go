package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	for _, k := range arg {
		if k == "01" || k == "galaxy" || k == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		} else {
		}
	}
}

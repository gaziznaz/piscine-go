package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args
	len := 0
	for i := range arg {
		len = i
	}

	if len == 0 {
		fmt.Println("File name missing")
	} else if len > 1 {
		fmt.Println("Too many arguments")
	} else {

		fContent, err := ioutil.ReadFile(arg[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf(string(fContent))
	}
}

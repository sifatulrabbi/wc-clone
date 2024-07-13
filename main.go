package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		args     = os.Args[1:]
		options  = []string{}
		filename = ""
	)
	for i := 0; i < len(args); i++ {
		input := args[i]
		if string(input[0]) == "-" {
			options = append(options, strings.Split(input, "")[1:]...)
		} else {
			filename = input
		}
	}
	if filename == "" {
		log.Panicln("Please input the targeted file name!")
	}

	fmt.Printf("targeted file: '%s'\n", filename)
	fmt.Printf("options: '%v'\n", options)

	f, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Panicln("Selected file does not exists.", err)
	} else if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("targeted file size %d bytes\n", f.Size())
}

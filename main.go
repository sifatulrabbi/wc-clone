package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		args     = os.Args[1:]
		options  = map[string]string{}
		filename = ""
	)
	for i := 0; i < len(args); i++ {
		input := args[i]
		if string(input[0]) == "-" {
			for _, op := range strings.Split(input, "")[1:] {
				options[op] = op
			}
		} else {
			filename = input
		}
	}
	if filename == "" {
		log.Panicln("Please input the targeted file name!")
	}

	// fmt.Printf("targeted file: '%s'\n", filename)
	// fmt.Printf("options: '%v'\n", options)

	f, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Panicln("Selected file does not exists.", err)
	} else if err != nil {
		log.Panicln(err)
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Panicln(err)
	}

	// bytes -c --bytes
	// words -w --words
	// chars -m --chars
	// lines -l --lines
	// max line length -L --max-line-length
	// --help

	var (
		totalBytes    = f.Size()
		totalChars    = 0
		totalWords    = 0
		totalLines    = 0
		maxLineLength = 0
		result        = "file: " + filename
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalLines++
		totalWords += strings.Count(line, " ")
		totalChars += len(line)
		if maxLineLength < len(line) {
			maxLineLength = len(line)
		}
	}

	result += fmt.Sprintf("\nlines: %d", totalLines)
	for op := range options {
		if op == "m" {
			result += fmt.Sprintf("\nchars: %d", totalChars)
		}
		if op == "w" {
			result += fmt.Sprintf("\nwords: %d", totalWords)
		}
		if op == "c" {
			result += fmt.Sprintf("\nbytes: %d", totalBytes)
		}
		if op == "L" {
			result += fmt.Sprintf("\nMax line length: %d", totalBytes)
		}
	}
	fmt.Println(result)
}

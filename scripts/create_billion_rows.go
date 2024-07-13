package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	filename      string = "one-billion-rows.txt"
	chars         string = "aA0bB1cC2dD3eE4fF5gG6hH7iI8jJ9kK0lL1mM2nN3oO4pP5qQ6rR7sS8tT9uU0vVwWxXyYzZ"
	targetedLines int    = 1_000_000_000
	// targetedLines int    = 1_000
)

type TestFile struct {
	LinesCount int
	target     *os.File
	mu         *sync.Mutex
	wg         *sync.WaitGroup
	q          chan string
	errors     []error
}

func NewTestFile(f *os.File) *TestFile {
	file := TestFile{
		LinesCount: 0,
		target:     f,
		q:          make(chan string, 100_000),
		mu:         &sync.Mutex{},
		wg:         &sync.WaitGroup{},
		errors:     []error{},
	}
	go file.PrepareWriters(100_000)
	return &file
}

func (f *TestFile) Write(str string) {
	f.wg.Add(1)
	f.q <- str
}

func (f *TestFile) PrepareWriters(count int) {
	for i := 0; i < count; i++ {
		go func() {
			for str := range f.q {

				f.mu.Lock()
				if _, err := f.target.WriteString(str); err != nil {
					f.errors = append(f.errors, err)
					log.Panicln(err)
				}
				f.LinesCount++
				f.mu.Unlock()
				f.wg.Done()
			}
		}()
	}
}

func (f *TestFile) Wait() {
	f.wg.Wait()
	close(f.q)
	if err := f.target.Close(); err != nil {
		f.errors = append(f.errors, err)
	}
	if len(f.errors) > 0 {
		log.Fatalln(len(f.errors), "errors occurred during the file creation")
	} else {
		fmt.Println("done creating the file")
	}
}

func getLine() string {
	l := []byte{}
	for i := 0; i < 8; i++ {
		r := rand.Intn(len(chars))
		l = append(l, chars[r])
	}
	return string(l)
}

func main() {
	fmt.Println("Creating a file with 1 billion rows. Required free space on disk is: 20GB.")
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		if err = os.Remove(filename); err != nil {
			log.Panicln(err)
		}
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
	}
	file := NewTestFile(f)
	start := time.Now()

	for i := 0; i < targetedLines; i++ {
		file.Write(fmt.Sprintf("%s %s %s %s %s %s %s %s\n",
			getLine(), getLine(), getLine(), getLine(), getLine(), getLine(), getLine(), getLine()))
	}

	// for i := 0; i < 100_000; i++ {
	// 	go func() {
	// 		for file.LinesCount < targetedLines {
	// 			file.Write(fmt.Sprintf("%s %s %s %s %s %s %s %s\n",
	// 				getLine(), getLine(), getLine(), getLine(), getLine(), getLine(), getLine(), getLine()))
	// 		}
	// 	}()
	// }

	file.Wait()
	fmt.Printf("Execution took: %s\n", time.Since(start).String())
}

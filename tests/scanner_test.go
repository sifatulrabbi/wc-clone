package tests

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestScanningLines(t *testing.T) {
	f, err := os.OpenFile("../one-thousand-rows.txt", os.O_RDONLY, 0644)
	// f, err := os.OpenFile("../one-billion-rows.txt", os.O_RDONLY, 0644)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	sc := bufio.NewScanner(f)
	lines := 0
	start := time.Now()

	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, 1024)
	// sc.Scan()
	// fmt.Println(string(buf))

	for sc.Scan() {
		lines++
	}

	fmt.Println(time.Since(start).String())
	fmt.Printf("total lines: %d\n", lines)
}

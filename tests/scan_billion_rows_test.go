package tests

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"testing"
)

type Count struct {
	Val int
	mu  *sync.Mutex
}

func (c *Count) Increase(n int) {
	c.mu.Lock()
	c.Val += n
	c.mu.Unlock()
}

func TestScanningBillionLines(t *testing.T) {
	f, err := os.OpenFile("../one-billion-rows.txt", os.O_RDONLY, 0644)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	count := Count{Val: 0, mu: &sync.Mutex{}}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		count.Increase(1)
	}

	fmt.Println("Total lines:", count.Val)
}

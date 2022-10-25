package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup, counter int) {
	defer group.Done()
	var mutex sync.Mutex

	mutex.Lock()
	group.Add(1)

	fmt.Println("Hello", counter)

	mutex.Unlock()
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group, i)
	}

	group.Wait()
	fmt.Println("done")
}

package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	group := &sync.WaitGroup{}

	var pool sync.Pool

	// Make Default Value Pool
	pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("qwe")
	pool.Put("asd")
	pool.Put("zxc")

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Wait()

	// time.Sleep(3 * time.Second)
}

package goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("total CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total Goroutine", totalGoroutine)

}

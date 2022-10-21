package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	// close untuk menutup channel
	// recommend use defer
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "tes"
		fmt.Println("D.O.N.E")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeRespone(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "QWE"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeRespone(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(3 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "QWE"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Test"
	channel <- "Test 01"
	channel <- "Test 003"

	fmt.Println(<-channel)
	fmt.Println("D.O.N.E")
}

func TestRangeChannel(t *testing.T) {

	channel := make(chan string)

	go func() {
		for i := 0; i < 100000; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("data diterima", data)
	}

	fmt.Println("DONE")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespone(channel1)
	go GiveMeRespone(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2 ", data)
			counter++
		default:
			fmt.Println("waiting")
		}

		if counter == 2 {
			break
		}
	}
}

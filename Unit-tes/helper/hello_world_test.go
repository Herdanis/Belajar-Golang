package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("yes")
	if result != "Hello yes" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorld Done FailNow")
}
func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("yes")
	if result != "Hello yes" {
		t.Fail()
	}
	fmt.Println("TestHelloWorld Done Fail")
}

func TestHelloWorld3(t *testing.T) {
	result := HelloWorld("yes")
	if result != "Hello yes" {
		t.Error("Result Not Valid")
	}
	fmt.Println("TestHelloWorld Done Error")
}
func TestHelloWorld4(t *testing.T) {
	result := HelloWorld("yes")
	if result != "Hello yes" {
		t.Fatal("Result Not Valid")
	}
	fmt.Println("TestHelloWorld Done Fatal")
}

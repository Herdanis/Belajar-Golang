package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Tes", func(t *testing.T) {
		result := HelloWorld("yes")
		require.Equal(t, "Hello yes", result, "TestHelloWorld SubTest")
	})
}
func TestMain(m *testing.M) {
	fmt.Println("Before unit Test")

	m.Run()

	fmt.Println("After unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cant run on MacOS")
	}

	result := HelloWorld("tes")
	require.Equal(t, "Hello yes", result, "TestHelloWorld Done Skip")
}

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

func TestHelloWorld5(t *testing.T) {
	result := HelloWorld("yes")
	assert.Equal(t, "Hello yes", result, "Result Not Valid")
	fmt.Println("TestHelloWorld Done Assert Done")
}

func TestHelloWorld6(t *testing.T) {
	result := HelloWorld("yes")
	require.Equal(t, "Hello yes", result, "Result Not Valid")
	fmt.Println("TestHelloWorld Done Requere Done")
}

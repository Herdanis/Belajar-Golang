package main_test

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultiFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var filesP embed.FS

func TestPatchMatch(t *testing.T) {
	dirEntries, _ := filesP.ReadDir("files")
	for _, dir := range dirEntries {
		if !dir.IsDir() {
			fmt.Println(dir.Name())
			file, _ := filesP.ReadFile("files/" + dir.Name())
			fmt.Println(string(file))
		}
	}
}

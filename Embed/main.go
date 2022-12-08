package main

import (
	"embed"
	"fmt"
)

//go:embed version.txt
var version string

//go:embed files/*.txt
var files embed.FS

func main() {
	fmt.Println(version)

	dirEntries, _ := files.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := files.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}

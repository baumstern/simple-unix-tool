package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	path := "."
	if 1 < len(os.Args) {
		path = os.Args[1]
	}

	fileSystem := os.DirFS(path)

	dirEntry, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		log.Panic(err)
	}
	for _, d := range dirEntry {
		fmt.Println(d.Name())
	}
}

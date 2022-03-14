package main

import (
	"os"
)

func main() {
	data, err := os.ReadFile("/dev/stdin")
	if err != nil {
		panic(err.Error())
	}
	os.Stdout.Write(data)
}

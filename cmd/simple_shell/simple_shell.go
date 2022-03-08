package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Shell struct {
	buf string
}

func NewShell() *Shell {
	return &Shell{}
}

func (s *Shell) getInput() {
	fmt.Print("$$ ")
	fmt.Fscanln(os.Stdin, &s.buf)
}

func (s *Shell) Exec() {
	cmd := exec.Command(s.buf)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println("err:", err)
	}
}

func main() {
	s := NewShell()

	for {
		s.getInput()
		s.Exec()
	}
}

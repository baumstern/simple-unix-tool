package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
	buf string
}

func NewShell() *Shell {
	return &Shell{}
}

func (s *Shell) getInput() {
	fmt.Print("$$ ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s.buf = scanner.Text()
}

func (s *Shell) Exec() {
	ss := strings.Split(s.buf, " ")
	cmd := exec.Command(ss[0], ss[1:]...)
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

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func atoi(a string) int {
	i, e := strconv.Atoi(a)
	if e != nil {
		panic(e)
	}
	return i
}

type move interface {
	Apply(s *state)
}

type spin struct {
	n int
}

func (m *spin) Apply(s *state) {
	s.Spin(m.n)
}

type exchange struct {
	a, b int
}

func (m *exchange) Apply(s *state) {
	s.Exchange(m.a, m.b)
}

type partner struct {
	a, b byte
}

func (m *partner) Apply(s *state) {
	s.Partner(m.a, m.b)
}

func newMove(str string) move {
	switch str[0] {
	case 's':
		return &spin{atoi(str[1:])}
	case 'x':
		args := strings.Split(str[1:], "/")
		return &exchange{atoi(args[0]), atoi(args[1])}
	case 'p':
		args := strings.Split(str[1:], "/")
		return &partner{byte(args[0][0]), byte(args[1][0])}
	default:
		panic(str)
	}
}

func parse(filename string) []move {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(bytes), ",")
	moves := make([]move, len(parts))

	for i, part := range parts {
		moves[i] = newMove(part)
	}

	return moves
}

func main() {
	moves := parse("input.txt")

	s := &state{[]byte("abcdefghijklmnop"), 0}

	// for i := 0; i < 1000*1000*1000; i++ {
	// 	for _, move := range moves {
	// 		move.Apply(s)
	// 	}
	// 	if s.String() == "abcdefghijklmnop" {
	// 		fmt.Println("Loop found after", i+1, "iterations")
	// 		return
	// 	}
	// }

	runs := 1000 * 1000 * 1000 % 60
	for i := 0; i < runs; i++ {
		for _, move := range moves {
			move.Apply(s)
		}
	}

	fmt.Println(s)
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type move struct {
	code string
	op1  string
	op2  string
}

func atoi(a string) int {
	i, e := strconv.Atoi(a)
	if e != nil {
		panic(e)
	}
	return i
}

func newMove(str string) move {
	switch str[0] {
	case 's':
		return move{code: "s", op1: str[1:]}
	case 'x':
		ops := strings.Split(str[1:], "/")
		return move{code: "x", op1: ops[0], op2: ops[1]}
	case 'p':
		ops := strings.Split(str[1:], "/")
		return move{code: "p", op1: ops[0], op2: ops[1]}
	default:
		panic(str)
	}
}

func spin(state string, n int) string {
	return state[len(state)-n:] + state[:len(state)-n]
}

func exchange(state string, a, b int) string {
	if a > b {
		a, b = b, a
	}
	return state[:a] + state[b:b+1] + state[a+1:b] + state[a:a+1] + state[b+1:]
}

func partner(state, a, b string) string {
	ai := strings.Index(state, a)
	bi := strings.Index(state, b)
	return exchange(state, ai, bi)
}

func (m move) Apply(state string) string {
	switch m.code {
	case "s":
		return spin(state, atoi(m.op1))
	case "x":
		return exchange(state, atoi(m.op1), atoi(m.op2))
	case "p":
		return partner(state, m.op1, m.op2)
	default:
		panic(m.code)
	}
}

func (m move) String() string {
	res := m.code + m.op1
	if m.op2 != "" {
		res += "/" + m.op2
	}
	return res
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
	// state := "abcde"
	// state = spin(state, 1)
	// fmt.Println(state)
	// state = exchange(state, 3, 4)
	// fmt.Println(state)
	// state = partner(state, "e", "b")
	// fmt.Println(state)

	moves := parse("input.txt")

	state := "abcdefghijklmnop"

	for _, move := range moves {
		state = move.Apply(state)
	}

	fmt.Println(state)
}

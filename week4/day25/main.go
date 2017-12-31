package main

import "fmt"

type value bool
type direction int
type state string

type action struct {
	write value
	move  direction
	next  state
}

type rule map[value]action
type ruleset map[state]rule

func main() {
	tape := map[int]value{}
	state := state("A")
	cursor := 0
	rules := realRuleset()

	iters := 12425180

	for i := 0; i < iters; i++ {
		val := tape[cursor]
		action := rules[state][val]

		tape[cursor] = action.write
		cursor += int(action.move)
		state = action.next
	}

	count := 0
	for _, val := range tape {
		if val {
			count++
		}
	}

	fmt.Println(count)
}

func testRuleset() ruleset {
	return ruleset{
		"A": rule{
			false: action{
				write: true,
				move:  1,
				next:  "B",
			},
			true: action{
				write: false,
				move:  -1,
				next:  "B",
			},
		},
		"B": rule{
			false: action{
				write: true,
				move:  -1,
				next:  "A",
			},
			true: action{
				write: true,
				move:  1,
				next:  "A",
			},
		},
	}
}

func realRuleset() ruleset {
	return ruleset{
		"A": rule{
			false: action{
				write: true,
				move:  1,
				next:  "B",
			},
			true: action{
				write: false,
				move:  1,
				next:  "F",
			},
		},
		"B": rule{
			false: action{
				write: false,
				move:  -1,
				next:  "B",
			},
			true: action{
				write: true,
				move:  -1,
				next:  "C",
			},
		},
		"C": rule{
			false: action{
				write: true,
				move:  -1,
				next:  "D",
			},
			true: action{
				write: false,
				move:  1,
				next:  "C",
			},
		},
		"D": rule{
			false: action{
				write: true,
				move:  -1,
				next:  "E",
			},
			true: action{
				write: true,
				move:  1,
				next:  "A",
			},
		},
		"E": rule{
			false: action{
				write: true,
				move:  -1,
				next:  "F",
			},
			true: action{
				write: false,
				move:  -1,
				next:  "D",
			},
		},
		"F": rule{
			false: action{
				write: true,
				move:  1,
				next:  "A",
			},
			true: action{
				write: false,
				move:  -1,
				next:  "E",
			},
		},
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type condition struct {
	Register   string // the register to compare
	Comparator string // the comparision function (<, >, ==, <=, >=, !=)
	Value      int    // the value to compare against
}

type operation struct {
	Register  string    // the register to update
	Opcode    string    // inc or dec
	Operand   int       // how much to increment or decrement by
	Condition condition // the condition to apply
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func parse(file string) []operation {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	ops := []operation{}

	for scan.Scan() {
		parts := strings.Split(scan.Text(), " ")
		ops = append(ops, operation{
			Register: parts[0],
			Opcode:   parts[1],
			Operand:  atoi(parts[2]),
			Condition: condition{
				Register:   parts[4],
				Comparator: parts[5],
				Value:      atoi(parts[6]),
			},
		})
	}

	if scan.Err() != nil {
		panic(scan.Err())
	}

	return ops
}

func eval(registers map[string]int, cond condition) bool {
	val := registers[cond.Register]
	switch cond.Comparator {
	case "<":
		return val < cond.Value
	case "<=":
		return val <= cond.Value
	case "==":
		return val == cond.Value
	case ">=":
		return val >= cond.Value
	case ">":
		return val > cond.Value
	case "!=":
		return val != cond.Value
	default:
		panic(cond.Comparator)
	}
}

func main() {
	operations := parse("input.txt")
	registers := map[string]int{}

	maxmax := 0

	for _, op := range operations {
		if eval(registers, op.Condition) {
			switch op.Opcode {
			case "inc":
				registers[op.Register] += op.Operand
			case "dec":
				registers[op.Register] -= op.Operand
			default:
				panic(op.Opcode)
			}

			if maxmax < registers[op.Register] {
				maxmax = registers[op.Register]
			}
		}
	}

	max := 0
	for _, v := range registers {
		if max < v {
			max = v
		}
	}

	fmt.Println(max)
	fmt.Println(maxmax)
}

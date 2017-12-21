package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	opcode string
	args   []string
}

func (i instruction) String() string {
	return fmt.Sprintf("%v %v", i.opcode, i.args)
}

func parse(filename string) []instruction {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	result := []instruction{}

	for scan.Scan() {
		parts := strings.Split(scan.Text(), " ")
		result = append(result, instruction{parts[0], parts[1:]})
	}

	if scan.Err() != nil {
		panic(scan.Err())
	}

	return result
}

type thread struct {
	id        int
	code      []instruction
	ip        int
	registers map[string]int
	input     <-chan int
	output    chan<- int
	sends     int
}

func newThread(code []instruction, id int, input <-chan int, output chan<- int) *thread {
	return &thread{
		id:        id,
		code:      code,
		ip:        0,
		registers: map[string]int{"p": id},
		input:     input,
		output:    output,
	}
}

func (t *thread) eval(str string) int {
	i, err := strconv.Atoi(str)
	if err == nil {
		return i
	}
	return t.registers[str]
}

func (t *thread) step() bool {
	op := t.code[t.ip]
	opcode := op.opcode
	args := op.args

	switch opcode {
	case "snd":
		val := t.eval(args[0])
		fmt.Printf("%v: snd %v (%v)\n", t.id, args[0], val)
		t.output <- val
		t.sends++
		t.ip++

	case "set":
		val := t.eval(args[1])
		t.registers[args[0]] = val
		fmt.Printf("%v: set %v %v (%v)\n", t.id, args[0], args[1], val)
		t.ip++

	case "add":
		val := t.eval(args[1])
		result := t.registers[args[0]] + val
		fmt.Printf("%v: add %v (%v) %v (%v) = %v\n", t.id, args[0], t.registers[args[0]], args[1], val, result)
		t.registers[args[0]] = result
		t.ip++

	case "mul":
		val := t.eval(args[1])
		result := t.registers[args[0]] * val
		fmt.Printf("%v: mul %v (%v) %v (%v) = %v\n", t.id, args[0], t.registers[args[0]], args[1], val, result)
		t.registers[args[0]] = result
		t.ip++

	case "mod":
		val := t.eval(args[1])
		result := t.registers[args[0]] % val
		fmt.Printf("%v: mod %v (%v) %v (%v) = %v\n", t.id, args[0], t.registers[args[0]], args[1], val, result)
		t.registers[args[0]] = result
		t.ip++

	case "rcv":
		select {
		case val := <-t.input:
			fmt.Printf("%v: rcv %v (%v)\n", t.id, args[0], val)
			t.registers[args[0]] = val
			t.ip++

		default:
			return false
		}

	case "jgz":
		x := t.eval(args[0])
		y := 1
		if x > 0 {
			y = t.eval(args[1])
		}
		fmt.Printf("%v: jgz %v (%v) %v (%v)\n", t.id, op.args[0], x, op.args[1], y)
		t.ip += y

	default:
		panic(op)
	}

	return true
}

func main() {
	code := parse("input.txt")

	in := make(chan int, 100000)
	out := make(chan int, 100000)

	t0 := newThread(code, 0, in, out)
	t1 := newThread(code, 1, out, in)

	for {
		ok0 := t0.step()
		ok1 := t1.step()

		if !ok0 && !ok1 {
			fmt.Println("deadlock:", t1.sends)
			break
		}
	}
}

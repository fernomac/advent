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

func eval(str string, registers map[string]int) int {
	i, err := strconv.Atoi(str)
	if err == nil {
		return i
	}
	return registers[str]
}

func exec(ip int, op instruction, registers map[string]int) int {
	switch op.opcode {
	case "snd":
		val := eval(op.args[0], registers)
		fmt.Printf("snd %v (%v)\n", op.args[0], val)
		registers["snd"] = val
		return ip + 1

	case "set":
		val := eval(op.args[1], registers)
		registers[op.args[0]] = val
		fmt.Printf("set %v %v (%v)\n", op.args[0], op.args[1], val)
		return ip + 1

	case "add":
		val := eval(op.args[1], registers)
		result := registers[op.args[0]] + val
		fmt.Printf("add %v (%v) %v (%v) = %v\n", op.args[0], registers[op.args[0]], op.args[1], val, result)
		registers[op.args[0]] = result
		return ip + 1

	case "mul":
		val := eval(op.args[1], registers)
		result := registers[op.args[0]] * val
		fmt.Printf("mul %v (%v) %v (%v) = %v\n", op.args[0], registers[op.args[0]], op.args[1], val, result)
		registers[op.args[0]] = result
		return ip + 1

	case "mod":
		val := eval(op.args[1], registers)
		result := registers[op.args[0]] % val
		fmt.Printf("mod %v (%v) %v (%v) = %v\n", op.args[0], registers[op.args[0]], op.args[1], val, result)
		registers[op.args[0]] = result
		return ip + 1

	case "rcv":
		x := eval(op.args[0], registers)
		if x == 0 {
			// nop
			return ip + 1
		}
		panic(fmt.Sprintf("rcv %v (%v) = %v\n", op.args[0], x, registers["snd"]))

	case "jgz":
		x := eval(op.args[0], registers)
		if x > 0 {
			y := eval(op.args[1], registers)
			fmt.Printf("jgz %v (%v) %v (%v)\n", op.args[0], x, op.args[1], y)
			return ip + y
		}
		fmt.Printf("jgz %v (%v) %v (nop)\n", op.args[0], x, op.args[1])
		return ip + 1

	default:
		panic(op)
	}
}

func main() {
	code := parse("input.txt")
	// for _, op := range code {
	// 	fmt.Println(op)
	// }

	ip := 0
	registers := make(map[string]int)

	for {
		if ip < 0 || ip >= len(code) {
			break
		}
		op := code[ip]
		ip = exec(ip, op, registers)
	}
}

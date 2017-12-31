package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instr struct {
	op, x, y string
}

type proc struct {
	primes []int
	code   []instr
	ip     int
	reg    map[string]int
}

func (p *proc) Eval(val string) int {
	i, err := strconv.Atoi(val)
	if err == nil {
		return i
	}
	return p.reg[val]
}

func primesUpTo(n int) []int {
	primes := []int{2}

	for i := 3; i < n; i += 2 {
		prime := true

		for _, t := range primes {
			if i%t == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes = append(primes, i)
		}
	}

	return primes
}

func (p *proc) isPrime(n int) bool {
	if p.primes == nil {
		p.primes = primesUpTo(500)
	}

	for _, prime := range p.primes {
		if n%prime == 0 {
			return false
		}
	}

	return true
}

func (p *proc) Run() {
	for {
		if p.ip < 0 || p.ip >= len(p.code) {
			return
		}

		if p.ip == 8 {
			b := p.reg["b"]

			// f gets set to 1 if b is prime.
			if p.isPrime(b) {
				fmt.Println("found prime:", b)
				p.reg["f"] = 1
			} else {
				p.reg["f"] = 0
			}

			p.ip = 24 // jump outside of the outer loop.
		}

		instr := p.code[p.ip]

		switch instr.op {
		case "set":
			y := p.Eval(instr.y)
			// fmt.Printf("%v:\tset %v %v (%v)\n", p.ip, instr.x, instr.y, y)
			p.reg[instr.x] = y
			p.ip++

		case "sub":
			x := p.reg[instr.x]
			y := p.Eval(instr.y)
			res := x - y
			// fmt.Printf("%v:\tsub %v (%v) %v (%v) = %v\n", p.ip, instr.x, x, instr.y, y, res)
			p.reg[instr.x] = res
			p.ip++

		case "mul":
			x := p.reg[instr.x]
			y := p.Eval(instr.y)
			res := x * y
			// fmt.Printf("%v:\tmul %v (%v) %v (%v) = %v\n", p.ip, instr.x, x, instr.y, y, res)
			p.reg[instr.x] = res
			p.ip++

		case "jnz":
			x := p.Eval(instr.x)
			y := p.Eval(instr.y)
			// fmt.Printf("%v:\tjnz %v (%v) %v (%v)\n", p.ip, instr.x, x, instr.y, y)
			if x == 0 {
				p.ip++
			} else {
				p.ip += y
			}

		default:
			panic("unknown op: " + instr.op)
		}
	}
}

func parse(filename string) []instr {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	result := []instr{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			panic("weird line: " + line)
		}

		result = append(result, instr{
			op: parts[0],
			x:  parts[1],
			y:  parts[2],
		})
	}

	return result
}

func main() {
	code := parse("input.txt")
	proc := proc{code: code, ip: 0, reg: make(map[string]int)}

	proc.ip = 0
	proc.reg = map[string]int{"a": 1}

	proc.Run()
	fmt.Println(proc.Eval("h"))
}

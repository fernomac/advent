package main

import (
	"bufio"
	"fmt"
	"os"
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

func parse(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	result := make([]int, 0)

	for scan.Scan() {
		parts := strings.Split(scan.Text(), ": ")
		idx := atoi(parts[0])
		for idx > len(result) {
			result = append(result, 0)
		}
		result = append(result, atoi(parts[1]))
	}

	if scan.Err() != nil {
		panic(scan.Err())
	}

	return result
}

type firewall struct {
	ranges    []int
	positions []int
	movingUps []bool
}

func newFirewall(ranges []int) *firewall {
	return &firewall{
		ranges:    ranges,
		positions: make([]int, len(ranges)),
		movingUps: make([]bool, len(ranges)),
	}
}

func (f *firewall) Step() {
	for i := 0; i < len(f.ranges); i++ {
		if f.ranges[i] > 0 {
			if f.movingUps[i] {
				f.positions[i]--
				if f.positions[i] == 0 {
					f.movingUps[i] = false
				}
			} else {
				f.positions[i]++
				if f.positions[i] == f.ranges[i]-1 {
					f.movingUps[i] = true
				}
			}
		}
	}
}

func (f *firewall) String() string {
	result := ""
	for i := 0; i < len(f.ranges); i++ {
		result += strconv.Itoa(i)

		if f.ranges[i] == 0 {
			result += "\t...\n"
		} else {
			for j := 0; j < f.ranges[i]; j++ {
				if j == f.positions[i] {
					result += "\t[S]"
				} else {
					result += "\t[ ]"
				}
			}
			result += "\n"
		}
	}
	return result
}

func run(ranges []int, delay int) (bool, int) {
	fw := newFirewall(ranges)
	for i := 0; i < delay; i++ {
		fw.Step()
	}

	depth := -1
	caught := false
	severity := 0

	for depth < len(ranges)-1 {
		depth++

		if ranges[depth] > 0 && fw.positions[depth] == 0 {
			caught = true
			severity += (depth * ranges[depth])
		}

		fw.Step()
	}

	return caught, severity
}

func mapit(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		result[i] = f(arr[i])
	}
	return result
}

func main() {
	// ranges := []int{3, 2, 0, 0, 4, 0, 4}
	ranges := parse("input.txt")

	modulos := mapit(ranges, func(val int) int {
		if val < 2 {
			return val
		}
		return ((val - 2) * 2) + 2
	})

	fmt.Println(modulos)

	delay := 0
	for {
		caught := false
		for i := 0; i < len(modulos); i++ {
			if modulos[i] > 0 {
				if (delay+i)%modulos[i] == 0 {
					caught = true
					break
				}
			}
		}
		if !caught {
			fmt.Println(delay)
			break
		}
		delay++
	}

	// delay := 0
	// for {
	// 	caught, severity := run(ranges, delay)
	// 	if delay == 0 {
	// 		fmt.Println("Severity with delay 0:", severity)
	// 	}
	// 	if !caught {
	// 		fmt.Println("Not caught with delay", delay)
	// 		break
	// 	}
	// 	if severity == 0 {
	// 		fmt.Println("Severity=0 at delay=", delay)
	// 	}
	// 	delay++
	// }
}

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

func main() {
	fw := newFirewall(parse("input.txt"))

	depth := -1
	severity := 0

	for depth < len(fw.ranges)-1 {
		depth++

		if fw.ranges[depth] > 0 && fw.positions[depth] == 0 {
			// Caught.
			severity += (depth * fw.ranges[depth])
		}

		// Move sensors.
		fw.Step()
	}

	fmt.Println(severity)
}

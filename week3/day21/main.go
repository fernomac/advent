package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type grid [][]bool

func (g grid) Size() int {
	return len(g)
}

func (g grid) BitsSet() int {
	count := 0
	for y := 0; y < g.Size(); y++ {
		for x := 0; x < g.Size(); x++ {
			if g[y][x] {
				count++
			}
		}
	}
	return count
}

func (g grid) Subgrid(x, y, size int) grid {
	result := grid{}

	for i := 0; i < size; i++ {
		result = append(result, g[y+i][x:x+size])
	}

	return result
}

func (g grid) Split() [][]grid {
	var mod int
	if g.Size()%2 == 0 {
		mod = 2
	} else if g.Size()%3 == 0 {
		mod = 3
	} else {
		panic(g.Size())
	}

	result := [][]grid{}

	for y := 0; y < g.Size(); y += mod {
		row := []grid{}

		for x := 0; x < g.Size(); x += mod {
			row = append(row, g.Subgrid(x, y, mod))
		}

		result = append(result, row)
	}

	return result
}

func (g grid) Rotate() grid {
	result := grid{}

	for x := 0; x < g.Size(); x++ {
		row := []bool{}
		for y := g.Size() - 1; y >= 0; y-- {
			row = append(row, g[y][x])
		}
		result = append(result, row)
	}

	return result
}

func (g grid) Flip() grid {
	result := grid{}

	for y := 0; y < g.Size(); y++ {
		row := []bool{}
		for x := g.Size() - 1; x >= 0; x-- {
			row = append(row, g[y][x])
		}
		result = append(result, row)
	}

	return result
}

func (g grid) Equals(o grid) bool {
	if g.Size() != o.Size() {
		return false
	}

	for x := 0; x < g.Size(); x++ {
		for y := 0; y < g.Size(); y++ {
			if g[y][x] != o[y][x] {
				return false
			}
		}
	}

	return true
}

func (g grid) String() string {
	buf := bytes.Buffer{}

	for _, row := range g {
		for _, bit := range row {
			if bit {
				buf.WriteByte('#')
			} else {
				buf.WriteByte('.')
			}
		}
		buf.WriteByte('\n')
	}

	return buf.String()
}

type rule struct {
	from, to grid
}

func (r *rule) String() string {
	buf := bytes.Buffer{}

	buf.WriteString(r.from.String())

	spaces := len(r.from) / 2
	for i := 0; i < spaces; i++ {
		buf.WriteByte(' ')
	}
	buf.WriteString("v\n")

	buf.WriteString(r.to.String())

	return buf.String()
}

func matches(state, from grid) bool {
	test := from
	for i := 0; i < 4; i++ {
		if state.Equals(test) {
			return true
		}
		test = test.Rotate()
	}

	test = from.Flip()
	for i := 0; i < 4; i++ {
		if state.Equals(test) {
			return true
		}
		test = test.Rotate()
	}

	return false
}

func transform(state grid, rules []rule) grid {
	for _, rule := range rules {
		if matches(state, rule.from) {
			return rule.to
		}
	}
	panic("no rules match:\n" + state.String())
}

func join(parts [][]grid) grid {
	result := grid{}

	for y := 0; y < len(parts); y++ {
		partrow := parts[y]
		for yy := 0; yy < partrow[0].Size(); yy++ {
			row := []bool{}
			for x := 0; x < len(partrow); x++ {
				row = append(row, partrow[x][yy]...)
			}
			result = append(result, row)
		}
	}

	return result
}

func main() {
	state := parseGrid(".#./..#/###")
	rules := parse("input.txt")

	for i := 0; i < 18; i++ {
		parts := state.Split()

		for y := 0; y < len(parts); y++ {
			for x := 0; x < len(parts); x++ {
				parts[y][x] = transform(parts[y][x], rules)
			}
		}

		state = join(parts)
		// fmt.Println(state)
	}

	fmt.Println(state.BitsSet())
}

func parseGrid(line string) grid {
	result := grid{}
	parts := strings.Split(line, "/")

	for _, part := range parts {
		if len(part) != len(parts) {
			panic("weird grid: " + line)
		}

		row := []bool{}

		for _, char := range part {
			if char == '#' {
				row = append(row, true)
			} else if char == '.' {
				row = append(row, false)
			} else {
				panic("weird char: " + line)
			}
		}

		result = append(result, row)
	}

	return result
}

func parseRule(line string) rule {
	parts := strings.Split(line, " => ")
	if len(parts) != 2 {
		panic("weird line: " + line)
	}

	from := parseGrid(parts[0])
	to := parseGrid(parts[1])

	return rule{from: from, to: to}
}

func parse(filename string) []rule {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	result := []rule{}

	for scan.Scan() {
		result = append(result, parseRule(scan.Text()))
	}

	if scan.Err() != nil {
		panic(scan.Err())
	}

	return result
}

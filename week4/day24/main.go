package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type component struct {
	a, b int
	used bool
}

func (c *component) Other(val int) int {
	if val == c.a {
		return c.b
	}
	if val == c.b {
		return c.a
	}
	panic("whoops")
}

func index(cs []*component) map[int][]*component {
	result := map[int][]*component{}

	for _, c := range cs {
		result[c.a] = append(result[c.a], c)
		result[c.b] = append(result[c.b], c)
	}

	return result
}

type score struct {
	length, strength int
}

func (s *score) IsGreaterThan(o *score) bool {
	if s.length > o.length {
		return true
	}
	return s.strength > o.strength
}

func maxBridge(parts map[int][]*component, start int) score {
	matches := parts[start]

	max := score{}

	for _, match := range matches {
		if !match.used {
			match.used = true

			other := match.Other(start)
			score := maxBridge(parts, other)
			score.length++
			score.strength += (start + other)

			if score.IsGreaterThan(&max) {
				max = score
			}

			match.used = false
		}
	}

	return max
}

func main() {
	cs := parse("input.txt")
	index := index(cs)

	max := maxBridge(index, 0)

	fmt.Println(max)
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func parse(filename string) []*component {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	result := []*component{}

	for _, line := range lines {
		parts := strings.Split(line, "/")
		if len(parts) != 2 {
			panic("weird line: " + line)
		}
		result = append(result, &component{
			a:    atoi(parts[0]),
			b:    atoi(parts[1]),
			used: false,
		})
	}

	return result
}

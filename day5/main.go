package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func load(file string) []int {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ret := []int{}

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		val, err := strconv.Atoi(scan.Text())
		if err != nil {
			panic(err)
		}
		ret = append(ret, val)
	}

	if scan.Err() != nil {
		panic(scan.Err())
	}

	return ret
}

func main() {
	maze := load("input.txt")

	pos := 0
	steps := 0

	for pos >= 0 && pos < len(maze) {
		offset := maze[pos]
		maze[pos] = maze[pos] + 1
		pos += offset
		steps++
	}

	fmt.Println(steps)
}

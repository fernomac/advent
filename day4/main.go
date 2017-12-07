package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValid(phrase []string) bool {
	lut := make(map[string]struct{})

	for _, word := range phrase {
		if _, ok := lut[word]; ok {
			return false
		}
		lut[word] = struct{}{}
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		phrase := strings.Split(scanner.Text(), " ")
		if isValid(phrase) {
			sum++
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(sum)
}

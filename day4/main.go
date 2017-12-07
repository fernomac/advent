package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type sorter []rune

func (s sorter) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sorter) Len() int {
	return len(s)
}

func isMoreValid(phrase []string) bool {
	lut := make(map[string]struct{})

	for _, word := range phrase {
		runes := []rune(word)
		sort.Sort(sorter(runes))
		if _, ok := lut[string(runes)]; ok {
			return false
		}
		lut[string(runes)] = struct{}{}
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
	sum2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		phrase := strings.Split(scanner.Text(), " ")
		if isValid(phrase) {
			sum++
		}
		if isMoreValid(phrase) {
			sum2++
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

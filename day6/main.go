package main

import "fmt"

func findmax(banks []int) int {
	index := 0
	value := banks[0]

	for i := 1; i < len(banks); i++ {
		if banks[i] > value {
			index = i
			value = banks[i]
		}
	}

	return index
}

func main() {
	banks := []int{2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14}
	seen := make(map[string]struct{})
	seen[fmt.Sprint(banks)] = struct{}{}

	rounds := 0
	for {
		index := findmax(banks)

		blocks := banks[index]
		banks[index] = 0

		for blocks > 0 {
			index++
			if index >= len(banks) {
				index = 0
			}
			banks[index]++
			blocks--
		}

		rounds++

		key := fmt.Sprint(banks)
		if _, ok := seen[key]; ok {
			break
		}
		seen[key] = struct{}{}
	}

	fmt.Println(rounds)
}

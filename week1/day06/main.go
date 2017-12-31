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
	seen := make(map[string]int)
	seen[fmt.Sprint(banks)] = 0

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
		if val, ok := seen[key]; ok {
			fmt.Println("finished at round", rounds)
			fmt.Println("cycle length", rounds-val)
			break
		}
		seen[key] = rounds
	}
}

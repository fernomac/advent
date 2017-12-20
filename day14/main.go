package main

import "fmt"
import "strconv"

func bitmask(data []byte, offset int) bool {
	b := data[offset/8]
	mask := byte(1 << uint(7-(offset%8)))
	return ((b & mask) == mask)
}

func usage(input string) int {
	used := 0
	for row := 0; row < 128; row++ {
		hash := knothash(input + "-" + strconv.Itoa(row))
		for column := 0; column < 128; column++ {
			if bitmask(hash, column) {
				used++
			}
		}
	}
	return used
}

func main() {
	fmt.Println(usage("flqrgnkx"))
	fmt.Println(usage("ljoxqyyw"))
}
